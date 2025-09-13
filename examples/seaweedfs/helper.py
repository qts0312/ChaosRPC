import json
import subprocess
import sys

def update_pod_config(json_file_path, namespace="seaweedfs", config_path="/root/chaos_config.json"):
    try:
        with open(json_file_path, 'r') as f:
            data = json.load(f)

        call_site = data.get("call_site", "")
        if ";" not in call_site:
            print("Error: The 'call_site' value does not contain a semicolon.")
            return

        pod_name, new_call_site = call_site.split(";", 1)
        data["call_site"] = new_call_site

        new_json_content = json.dumps(data, indent=4)

    except FileNotFoundError:
        print(f"Error: The file '{json_file_path}' was not found.")
        return
    except json.JSONDecodeError:
        print(f"Error: Could not decode JSON from '{json_file_path}'. Please check the file format.")
        return

    try:
        pods_output = subprocess.run(
            ["kubectl", "get", "pods", "-n", namespace, "-l", f"app.kubernetes.io/component={pod_name}"],
            capture_output=True,
            text=True,
            check=True
        )

        pod_lines = pods_output.stdout.strip().split('\n')[1:]
        if not pod_lines:
            print(f"No Pods found with app.kubernetes.io/component={pod_name} in namespace '{namespace}'.")
            return

        pod_names = [line.split()[0] for line in pod_lines]

    except subprocess.CalledProcessError as e:
        print(f"Error executing kubectl command: {e}")
        print(f"Standard Error: {e.stderr}")
        return

    for pod in pod_names:
        try:
            command = [
                "kubectl", "exec", "-n", namespace, pod, "--",
                "sh", "-c", f"echo '{new_json_content}' > {config_path}"
            ]

            subprocess.run(
                command,
                check=True,
                capture_output=True
            )

        except subprocess.CalledProcessError as e:
            print(f"Failed to update config for Pod {pod}: {e}")
            print(f"Standard Error: {e.stderr}")
            continue

if __name__ == "__main__":
    # Example usage:
    # Save your input JSON as a file, e.g., 'input.json'
    # Then run the script: python your_script_name.py input.json
    if len(sys.argv) < 2:
        print("Usage: python update_config.py <path_to_input_json_file>")
    else:
        input_json_file = sys.argv[1]
        update_pod_config(input_json_file)