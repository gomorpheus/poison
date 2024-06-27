# Findings

## Instance creation


This endpoint is hit: 

```json
/api/cmdb/v1.0/instances
```
Sends this request body:

```json

     {
         "class_name_key":{
            "name":"BMC_ComputerSystem",
            "namespace":"BMC.CORE"
         },
         "dataset_id":"BMC.ASSET",
         "attributes":{
            "Name":"linux-vmware 1-dev-All Clouds-ubuntu-104", // instance name
            "HostName":"linux-vmware-1-dev-all-clouds-ubuntu-104-.localdomain", // hostname
            "ShortDescription":"Provisioned from Morpheus: https://morpheuslab/provisioning/instances/277/nodes/126985\nTenant: Nexio\nUser: Admin User" // link to instance, tenant, and user who created 
         }
      }
   ]
}

```

## Instance/Hostname edit

CMDB not updated


## Instance deletion

Does not appear to delete the CMDB record