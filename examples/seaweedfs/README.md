# SeaweedFS

We provide reproducing scripts for some issues of [SeaweedFS](https://github.com/seaweedfs/seaweedfs) in this directory.

## Issues

Testcases are all in `ceph/s3-tests/s3tests_boto3/functional/test_s3.py`.

| Index                                                      | Description                                                               | Testcase                     |
|------------------------------------------------------------|---------------------------------------------------------------------------|------------------------------|
| [7203](https://github.com/seaweedfs/seaweedfs/issues/7203) | `CreateEntry` retries incorrectly on transient error when creating bucket | `test_bucket_list_empty`     |
| [7204](https://github.com/seaweedfs/seaweedfs/issues/7204) | `DeleteEntry` retries incorrectly on transient error                      | `test_bucket_list_empty`     |
| [7220](https://github.com/seaweedfs/seaweedfs/issues/7220) | Transient error on `LookupDirectoryEntry` leads to `NoSuchBucket` error   | `test_bucket_list_empty`     |
| [7221](https://github.com/seaweedfs/seaweedfs/issues/7221) | Transient error on `ListEntries` causes listing multipart uploads fail    | `test_list_multipart_upload` |
| [7224](https://github.com/seaweedfs/seaweedfs/issues/7224) | Transient error on `DeleteEntry` causes object deletion fail              | `test_multi_object_delete`   |
| [7228](https://github.com/seaweedfs/seaweedfs/issues/7228) | Timeout on `Assign` causes operation stuck for one minute                 | `test_bucket_list_many`      |
| [7229](https://github.com/seaweedfs/seaweedfs/issues/7229) | Timeout on `CollectionList` causes operation stuck for one minute         | `test_bucket_list_empty`     |

## Steps

1. Make tailored SeaweedFS image with ChaosRPC. Our tailored version is in [this fork](https://github.com/qts0312/seaweedfs).

   ```bash
   cd seaweedfs/docker
   make build
   ```

2. Deploy SeaweedFS in Kubernetes cluster with helm chart.

3. Use `helper.py` to modify configurations in pods to injects faults on specified call sites.

    ```bash
    python3 helper.py ./issue-<issue_index>.json
    ```
   
4. Run corresponding testcase.

## Notes

- To avoid endless injection in single testcase, we use `test_id` to identify each testcase run, and only inject faults once for each `test_id`. So If you want to rerun the testcase, please remember to change `test_id` in json file.
