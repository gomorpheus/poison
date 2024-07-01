# Findings

## Instance creation

### Computer System

```json
/api/cmdb/v1.0/instances POST
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_ComputerSystem",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Name": "op-test-remedy-3",
        "HostName": "op-test-remedy-3.localdomain",
        "ShortDescription": "Provisioned from Morpheus: https://morpheuslab/provisioning/instances/282/nodes/126990\nTenant: Neo\nUser: Admin User"
      }
    }
  ]
}
```


### Memory

```json
/api/cmdb/v1.0/instances POST
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_Memory",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Name": "op-test-remedy-3 (Memory)",
        "ShortDescription": "2048 MB"
      }
    }
  ]
}
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_Component",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Source.DatasetId": "BMC.ASSET",
        "ImpactDestinationId": "poisonInstance",
        "ClassId": "BMC.CORE:BMC_COMPONENT",
        "ImpactPropagationModel": "DIRECT",
        "Source.InstanceId": "poisonInstance",
        "Name": "BMC_Component",
        "ShortDescription": "na",
        "Destination.DatasetId": "BMC.ASSET",
        "Status": "0",
        "Destination.InstanceId": "poisonInstance",
        "HasImpact": "No",
        "ImpactWeight": 100,
        "DatasetId": "BMC.ASSET",
        "Destination.ClassId": "BMC_MEMORY",
        "ImpactDirection": "Unknown",
        "Source.ClassId": "BMC_COMPUTERSYSTEM",
        "ImpactSourceId": "poisonInstance"
      }
    }
  ]
}
```

### CPU 

```json
/api/cmdb/v1.0/instances POST
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_Processor",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Name": "op-test-remedy-3 (CPU)",
        "ShortDescription": "CPU count: 1"
      }
    }
  ]
}
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_Component",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Source.DatasetId": "BMC.ASSET",
        "ImpactDestinationId": "poisonInstance",
        "ClassId": "BMC.CORE:BMC_COMPONENT",
        "ImpactPropagationModel": "DIRECT",
        "Source.InstanceId": "poisonInstance",
        "Name": "BMC_Component",
        "ShortDescription": "na",
        "Destination.DatasetId": "BMC.ASSET",
        "Status": "0",
        "Destination.InstanceId": "poisonInstance",
        "HasImpact": "No",
        "ImpactWeight": 100,
        "DatasetId": "BMC.ASSET",
        "Destination.ClassId": "BMC_PROCESSOR",
        "ImpactDirection": "Unknown",
        "Source.ClassId": "BMC_COMPUTERSYSTEM",
        "ImpactSourceId": "poisonInstance"
      }
    }
  ]
}
```

### IP Address

 ```json
/api/cmdb/v1.0/instances POST
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_IPEndpoint",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Name": "10.32.23.62",
        "ShortDescription": "IP Address: 10.32.23.62"
      }
    }
  ]
}
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_Component",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Source.DatasetId": "BMC.ASSET",
        "ImpactDestinationId": "poisonInstance",
        "ClassId": "BMC.CORE:BMC_COMPONENT",
        "ImpactPropagationModel": "DIRECT",
        "Source.InstanceId": "poisonInstance",
        "Name": "BMC_Component",
        "ShortDescription": "na",
        "Destination.DatasetId": "BMC.ASSET",
        "Status": "0",
        "Destination.InstanceId": "poisonInstance",
        "HasImpact": "No",
        "ImpactWeight": 100,
        "DatasetId": "BMC.ASSET",
        "Destination.ClassId": "BMC_IPENDPOINT",
        "ImpactDirection": "Unknown",
        "Source.ClassId": "BMC_COMPUTERSYSTEM",
        "ImpactSourceId": "poisonInstance"
      }
    }
  ]
}
```

### Disk Drive

```json
/api/cmdb/v1.0/instances POST
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_DiskDrive",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Name": "op-test-remedy-3 (Disk root)",
        "ShortDescription": "Size: 20 GB",
        "Size": "21474836480"
      }
    }
  ]
}
```

```json
{
  "instances": [
    {
      "class_name_key": {
        "name": "BMC_Component",
        "namespace": "BMC.CORE"
      },
      "dataset_id": "BMC.ASSET",
      "attributes": {
        "Source.DatasetId": "BMC.ASSET",
        "ImpactDestinationId": "poisonInstance",
        "ClassId": "BMC.CORE:BMC_COMPONENT",
        "ImpactPropagationModel": "DIRECT",
        "Source.InstanceId": "poisonInstance",
        "Name": "BMC_Component",
        "ShortDescription": "na",
        "Destination.DatasetId": "BMC.ASSET",
        "Status": "0",
        "Destination.InstanceId": "poisonInstance",
        "HasImpact": "No",
        "ImpactWeight": 100,
        "DatasetId": "BMC.ASSET",
        "Destination.ClassId": "BMC_DISKDRIVE",
        "ImpactDirection": "Unknown",
        "Source.ClassId": "BMC_COMPUTERSYSTEM",
        "ImpactSourceId": "poisonInstance"
      }
    }
  ]
}
```


### Asset Life Cycle

/api/arsys/v1/entry/AST:ComputerSystem/{request_id}

```json
{
  "values": {
    "AssetLifecycleStatus": "Down",
    "Installation Date": "2024-07-01T07:45:52.000+0000"
  }
}
```

```json
{
  "values": {
    "AssetLifecycleStatus": "Deployed",
    "Installation Date": "2024-07-01T07:45:52.000+0000"
  }
}
```


