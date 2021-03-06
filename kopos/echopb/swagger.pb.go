package echopb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/boot": {
      "post": {
        "summary": "To boot virtual machines into networking landscape (App user, e.g. identity service register)",
        "description": "Input/Output is a same protobuf/json object. For example:\n {\n   \"flavor_name\": \"m1.small\",\n   \"image_name\": \"cirros\",\n   \"min_count\": 2,\n   \"max_count\": 4,\n   \"secgroups_info\": [],\n   \"user_data\": [],\n   \"network_name\": \"private\",\n   \"floating_network_name\": \"public\",\n   \"personality\": [],\n   \"name_prefix\": \"awesome VM\"\n }\nAltervative 'flavor_id' and 'image_id' is available",
        "operationId": "BootVirtualMachines",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackOpenstackNovaBootReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackOpenstackNovaBootReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/echo": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/flavors": {
      "get": {
        "summary": "To discover Flavors (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DiscoverFlavors",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackFlavorDiscoveryReqRespData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/flavors/0x3f/{id}": {
      "get": {
        "summary": "To discover a flavor detailed",
        "description": "The request input is 'id' in URL path",
        "operationId": "DiscoverFlavorDetailed",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbopenstackFlavor"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/flavors/0x3fname/{name}": {
      "get": {
        "summary": "To search flavor name for details",
        "description": "The request input is 'name' in URL path",
        "operationId": "SearchFlavorDetails",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbopenstackFlavor"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/images": {
      "get": {
        "summary": "To discover images (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DiscoverImages",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackImageDiscoveryReqRespData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/images/0x3f/{id}": {
      "get": {
        "summary": "To discover a image detailed",
        "description": "The request input is 'id' in URL path",
        "operationId": "DiscoverImageDetailed",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbopenstackImage"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/images/0x3fname/{name}": {
      "get": {
        "summary": "To search image name for details",
        "description": "The request input is 'name' in URL path",
        "operationId": "SearchImageDetails",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbopenstackImage"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/landscape": {
      "post": {
        "summary": "To establish networking landscape (App admin).",
        "description": "Input/Output is a same protobuf/json object but reqest only need less fields. For example:\n {\n   \"vnets\": [\n     {\n       \"name\": \"int-stage-0\",\n       \"subnets\": [\n         {\n           \"name\": \"int-192-168-128-0-slash-24\",\n           \"cidr\": \"192.168.128.0/24\",\n           \"enabledhcp\": true\n         }\n       ]\n     },\n     {\n       \"name\": \"int-stage-1\",\n       \"subnets\": [\n         {\n           \"name\": \"int-192-168-129-0-slash-24\",\n           \"cidr\": \"192.168.129.0/24\",\n           \"enabledhcp\": true\n         }\n       ]\n     },\n     {\n       \"name\": \"int-stage-2\",\n       \"subnets\": [\n         {\n           \"name\": \"int-192-168-130-0-slash-24\",\n           \"cidr\": \"192.168.130.0/24\",\n           \"enabledhcp\": true\n         }\n       ]\n     },\n     {\n       \"name\": \"int-stage-3\",\n       \"subnets\": [\n         {\n           \"name\": \"int-192-168-131-0-slash-24\",\n           \"cidr\": \"192.168.131.0/24\",\n           \"enabledhcp\": true\n         }\n       ]\n     },\n     {\n       \"name\": \"public\",\n       \"admin_state_up\": true,\n       \"subnets\": [\n         {\n           \"name\": \"10.100.151.0/24\",\n           \"cidr\": \"10.100.151.0/24\",\n           \"gateway_ip\": \"10.100.151.1\",\n           \"allocation_pools\": {\n             \"start\": \"10.100.151.50\",\n             \"end\": \"10.100.151.240\"\n           },\n           \"enabledhcp\": false\n         }\n       ],\n       \"shared\": true\n     }\n   ],\n   \"vrouter\": {\n     \"name\": \"hack\",\n     \"admint_state_up\": true\n   },\n   \"secgroup\": {\n     \"name\": \"hack\",\n     \"rules\": [\n       {\n         \"direction\": \"ingress\",\n         \"protocol\": \"tcp\"\n       },\n       {\n         \"direction\": \"ingress\",\n         \"protocol\": \"udp\"\n       },\n       {\n         \"direction\": \"ingress\",\n         \"protocol\": \"icmp\"\n       }\n     ]\n   },\n   \"ifaces_info\": [\n     {\n       \"router_name\": \"hack\",\n       \"network_name\": \"int-stage-0\",\n       \"subnet_name\": \"int-192-168-128-0-slash-24\",\n       \"secgroups_info\": [\n         {\n           \"name\": \"hack\"\n         }\n       ]\n     },\n     {\n       \"router_name\": \"hack\",\n       \"network_name\": \"int-stage-1\",\n       \"subnet_name\": \"int-192-168-129-0-slash-24\",\n       \"secgroups_info\": [\n         {\n           \"name\": \"hack\"\n         }\n       ]\n     },\n     {\n       \"router_name\": \"hack\",\n       \"network_name\": \"int-stage-2\",\n       \"subnet_name\": \"int-192-168-130-0-slash-24\",\n       \"secgroups_info\": [\n         {\n           \"name\": \"hack\"\n         }\n       ]\n     },\n     {\n       \"router_name\": \"hack\",\n       \"network_name\": \"int-stage-3\",\n       \"subnet_name\": \"int-192-168-131-0-slash-24\",\n       \"secgroups_info\": [\n         {\n           \"name\": \"hack\"\n         }\n       ]\n     }\n   ],\n   \"gateways_info\": [\n     {\n       \"network_name\": \"public\",\n       \"router_name\": \"hack\"\n     }\n   ]\n }\nAltervative '***_id' is available",
        "operationId": "EstablishNetworkLandscape",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackOpenstackNeutronLandscapeReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackOpenstackNeutronLandscapeReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/libvirt-domains0x3f/{server_id}": {
      "get": {
        "summary": "To discover Libvirt Domain info (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "GetLibvirtDomainVNCDisplay",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackLibvirtDomainReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/networks": {
      "get": {
        "summary": "To discover networks (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DiscoverNetworks",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackNetworkDiscoveryReqRespData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/reboot": {
      "post": {
        "summary": "To reboot Machines (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "RebootMachines",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackMachineRebootReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackMachineRebootReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/recycle": {
      "post": {
        "summary": "To discover Machines (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DestroyMachines",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackMachineDestroyReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackMachineDestroyReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/servers": {
      "get": {
        "summary": "To discover Machines (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DiscoverMachines",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackMachineDiscoveryReqRespData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/spawn": {
      "post": {
        "summary": "To spawn machines into networking landscape (App user, e.g. identity service register)",
        "description": "Input/Output is a same protobuf/json object. For example:\n[ {\n   \"flavor_name\": \"m1.small\",\n   \"image_name\": \"cirros\",\n   \"min_count\": 2,\n   \"max_count\": 4,\n   \"secgroups_info\": [],\n   \"user_data\": [],\n   \"network_name\": \"private\",\n   \"floating_network_name\": \"public\",\n   \"personality\": [],\n   \"name_prefix\": \"awesome VM\"\n} ]\nAltervative 'flavor_id' and 'image_id' is available\n'floating_network_name' is optional",
        "operationId": "SpawnMachines",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackMachineSpawnsReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackMachineSpawnsReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/ssh": {
      "post": {
        "operationId": "MockSSH",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbSSHReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbSSHReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/ssh/{cmd}": {
      "get": {
        "operationId": "MockSSH",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbSSHReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "cmd",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/sshpub": {
      "get": {
        "operationId": "MockSSH",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbSSHReqRespData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/subnets": {
      "get": {
        "summary": "To discover subnets (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DiscoverSubnets",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackSubnetDiscoveryReqRespData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/token": {
      "post": {
        "summary": "user (e.g. registered user) to verify if token is available",
        "operationId": "ValidateToken",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackTokenReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackTokenReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/topology": {
      "post": {
        "summary": "To discover networks (App user, e.g. identity service register).",
        "description": "Input/Output is a same protobuf/json object. But for HTTP request, you may not need anything.",
        "operationId": "DiscoverNetworkingTopology",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackNetworkTopologyReqRespData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackNetworkTopologyReqRespData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v2/battlefields": {
      "post": {
        "summary": "admin (e.g. head referee) to create battlefield",
        "operationId": "AdminSharedNetworkCreation",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackOpenstackNeutronNetResponseData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackOpenstackNeutronNetRequestData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v2/combatrooms": {
      "post": {
        "summary": "user (e.g. registered user) to establish combatroom",
        "operationId": "ApplyConsoleIntoDnatWithNetworkAndMachine",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openstackConsoleResourceResponseData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openstackConsoleResourceRequestData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    }
  },
  "definitions": {
    "echopbEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "echopbSSHReqRespData": {
      "type": "object",
      "properties": {
        "cmd": {
          "type": "string",
          "format": "string"
        },
        "env": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "result": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        }
      }
    },
    "echopbopenstackFlavor": {
      "type": "object",
      "properties": {
        "disk": {
          "type": "integer",
          "format": "int32",
          "title": "The Disk and RA\u003c fields provide a measure of storage space offered by the flavor, in GB and MB, respectively.\nDisk int json:\"disk\"\nRAM  int json:\"ram\""
        },
        "id": {
          "type": "string",
          "format": "string",
          "title": "The Id field contains the flavor's unique identifier.\nFor example, this identifier will be useful when specifying which hardware configuration to use for a new server instance.\nID string json:\"id\""
        },
        "is_public": {
          "type": "boolean",
          "format": "boolean",
          "title": "IsPublic indicates whether the flavor is public.\nIsPublic bool json:\"is_public\""
        },
        "name": {
          "type": "string",
          "format": "string",
          "title": "The Name field provides a human-readable moniker for the flavor.\nName       string  json:\"name\"\nRxTxFactor float64 json:\"rxtx_factor\""
        },
        "ram": {
          "type": "integer",
          "format": "int32"
        },
        "rxtx_factor": {
          "type": "number",
          "format": "double"
        },
        "swap": {
          "type": "integer",
          "format": "int32",
          "title": "Swap indicates how much space is reserved for swap.\nIf not provided, this field will be set to 0.\nSwap int json:\"swap\""
        },
        "vcpus": {
          "type": "integer",
          "format": "int32",
          "title": "VCPUs indicates how many (virtual) CPUs are available for this flavor.\nVCPUs int json:\"vcpus\""
        }
      },
      "description": "Flavor records represent (virtual) hardware configurations for server resources in a region."
    },
    "echopbopenstackGatewayInfo": {
      "type": "object",
      "properties": {
        "network_id": {
          "type": "string",
          "format": "string"
        },
        "network_name": {
          "type": "string",
          "format": "string"
        },
        "router_id": {
          "type": "string",
          "format": "string"
        },
        "router_name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "echopbopenstackImage": {
      "type": "object",
      "properties": {
        "checksum": {
          "type": "string",
          "format": "string",
          "title": "Checksum is the checksum of the data that's associated with the image"
        },
        "container_format": {
          "type": "string",
          "format": "string",
          "description": "ContainerFormat is the format of the container.\nValid values are ami, ari, aki, bare, and ovf."
        },
        "create_at": {
          "type": "string",
          "format": "string",
          "title": "CreatedAt is the date when the image has been created.\ngoogle.protobuf.Timestamp created_at = 16;"
        },
        "disk_format": {
          "type": "string",
          "format": "string",
          "description": "DiskFormat is the format of the disk.\nIf set, valid values are ami, ari, aki, vhd, vmdk, raw, qcow2, vdi, and iso."
        },
        "file": {
          "type": "string",
          "format": "string",
          "description": "File is the trailing path after the glance endpoint that represent the location\nof the image or the path to retrieve it."
        },
        "id": {
          "type": "string",
          "format": "string",
          "title": "ID is the image UUID"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "description": "Metadata is a set of metadata associated with the image.\nImage metadata allow for meaningfully define the image properties\nand tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html."
        },
        "min_disk": {
          "type": "integer",
          "format": "int32",
          "description": "MinDiskGigabytes is the amount of disk space in GB that is required to boot the image."
        },
        "min_ram": {
          "type": "integer",
          "format": "int32",
          "description": "MinRAMMegabytes [optional] is the amount of RAM in MB that is required to boot the image."
        },
        "name": {
          "type": "string",
          "format": "string",
          "description": "Name is the human-readable display name for the image."
        },
        "owner": {
          "type": "string",
          "format": "string",
          "description": "Owner is the tenant the image belongs to."
        },
        "properties": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "description": "Properties is a set of key-value pairs, if any, that are associated with the image."
        },
        "protected": {
          "type": "boolean",
          "format": "boolean",
          "description": "Protected is whether the image is deletable or not."
        },
        "schema": {
          "type": "string",
          "format": "string",
          "description": "Schema is the path to the JSON-schema that represent the image or image entity."
        },
        "size": {
          "type": "string",
          "format": "int64",
          "description": "SizeBytes is the size of the data that's associated with the image."
        },
        "status": {
          "type": "string",
          "format": "string",
          "title": "Status is the image status. It can be \"queued\" or \"active\"\nSee imageservice/v2/images/type.go"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "description": "Tags is a list of image tags. Tags are arbitrarily defined strings\nattached to an image."
        },
        "updated_at": {
          "type": "string",
          "format": "string",
          "title": "UpdatedAt is the date when the last change has been made to the image or it's properties.\ngoogle.protobuf.Timestamp updated_at = 17;"
        },
        "visibility": {
          "type": "string",
          "format": "string",
          "description": "Visibility defines who can see/use the image."
        }
      },
      "description": "Image model\nDoes not include the literal image data; just metadata.\nreturned by listing images, and by fetching a specific image."
    },
    "identityTenant": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "format": "string",
          "description": "Description is a human-readable explanation of this Tenant's purpose."
        },
        "enabled": {
          "type": "boolean",
          "format": "boolean",
          "description": "Enabled indicates whether or not a tenant is active."
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "ID is a unique identifier for this tenant."
        },
        "name": {
          "type": "string",
          "format": "string",
          "description": "Name is a friendlier user-facing name for this tenant."
        }
      },
      "description": "Tenant is a grouping of users in the identity service."
    },
    "identityToken": {
      "type": "object",
      "properties": {
        "expires_at": {
          "type": "string",
          "format": "string",
          "title": "ExpiresAt provides a timestamp in ISO 8601 format, indicating when the authentication token becomes invalid.\nAfter this point in time, future API requests made using this authentication token will respond with errors.\nEither the caller will need to reauthenticate manually, or more preferably, the caller should exploit automatic re-authentication.\nSee the AuthOptions structure for more details.\ngoogle.protobuf.Timestamp expires_at = 2;"
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "ID provides the primary means of identifying a user to the OpenStack API.\nOpenStack defines this field as an opaque value, so do not depend on its content.\nIt is safe, however, to compare for equality."
        },
        "tenant": {
          "$ref": "#/definitions/identityTenant",
          "description": "Tenant provides information about the tenant to which this token grants access."
        }
      },
      "description": "Token provides only the most basic information related to an authentication token."
    },
    "identityTokenCredentialsV2": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronAddressPair": {
      "type": "object",
      "properties": {
        "ip_address": {
          "type": "string",
          "format": "string"
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronAllocationPool": {
      "type": "object",
      "properties": {
        "end": {
          "type": "string",
          "format": "string"
        },
        "start": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronFloatingIP": {
      "type": "object",
      "properties": {
        "fixed_ip_address": {
          "type": "string",
          "format": "string",
          "description": "The specific IP address of the internal port which should be associated\nwith the floating IP."
        },
        "floating_ip_address": {
          "type": "string",
          "format": "string",
          "description": "Address of the floating IP on the external network."
        },
        "floating_network_id": {
          "type": "string",
          "format": "string",
          "description": "UUID of the external network where the floating IP is to be created."
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "Unique identifier for the floating IP instance."
        },
        "port_id": {
          "type": "string",
          "format": "string",
          "description": "UUID of the port on an internal network that is associated with the floating IP."
        },
        "router_id": {
          "type": "string",
          "format": "string",
          "title": "The ID of the router used for this Floating-IP"
        },
        "status": {
          "type": "string",
          "format": "string",
          "description": "The condition of the API resource."
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "description": "Owner of the floating IP. Only admin users can specify a tenant identifier\nother than its own."
        }
      },
      "description": "FloatingIP represents a floating IP resource. A floating IP is an external\nIP address that is mapped to an internal port and, optionally, a specific\nIP address on a private network. In other words, it enables access to an\ninstance on a private network from an external network. For this reason,\nfloating IPs can only be defined on networks where the 'router:external'\nattribute (provided by the external network extension) is set to True."
    },
    "neutronHostRoute": {
      "type": "object",
      "properties": {
        "destination_cidr": {
          "type": "string",
          "format": "string"
        },
        "next_hop": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronIP": {
      "type": "object",
      "properties": {
        "ip_address": {
          "type": "string",
          "format": "string"
        },
        "subnet_id": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronInterfaceInfo": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string",
          "description": "The UUID of the interface."
        },
        "port_id": {
          "type": "string",
          "format": "string",
          "description": "The ID of the port that is a part of the subnet."
        },
        "subnet_id": {
          "type": "string",
          "format": "string",
          "description": "The ID of the subnet which this interface is associated with."
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "description": "Owner of the interface."
        }
      },
      "description": "InterfaceInfo represents information about a particular router interface. As\nmentioned above, in order for a router to forward to a subnet, it needs an\ninterface."
    },
    "neutronPort": {
      "type": "object",
      "properties": {
        "admin_state_up": {
          "type": "boolean",
          "format": "boolean"
        },
        "allowed_address_pairs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronAddressPair"
          }
        },
        "device_id": {
          "type": "string",
          "format": "string"
        },
        "device_owner": {
          "type": "string",
          "format": "string"
        },
        "fixed_ips": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronIP"
          }
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "network_id": {
          "type": "string",
          "format": "string"
        },
        "security_groups": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "status": {
          "type": "string",
          "format": "string"
        },
        "tenant_id": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronRoute": {
      "type": "object",
      "properties": {
        "destination_cidr": {
          "type": "string",
          "format": "string"
        },
        "next_hop": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronRouter": {
      "type": "object",
      "properties": {
        "admin_state_up": {
          "type": "boolean",
          "format": "boolean"
        },
        "distributed": {
          "type": "boolean",
          "format": "boolean"
        },
        "gateway_info": {
          "$ref": "#/definitions/openstackneutronGatewayInfo"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronRoute"
          }
        },
        "status": {
          "type": "string",
          "format": "string"
        },
        "tenant_id": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "neutronSecGroup": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "format": "string",
          "description": "The security group description."
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "The UUID for the security group."
        },
        "name": {
          "type": "string",
          "format": "string",
          "description": "Human-readable name for the security group. Might not be unique. Cannot be\nnamed \"default\" as that is automatically created for a tenant."
        },
        "security_group_rules": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronSecGroupRule"
          },
          "description": "A slice of security group rules that dictate the permitted behaviour for\ntraffic entering and leaving the group."
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "description": "Owner of the security group. Only admin users can specify a TenantID\nother than their own."
        }
      },
      "description": "SecGroup represents a container for security group rules."
    },
    "neutronSecGroupRule": {
      "type": "object",
      "properties": {
        "direction": {
          "type": "string",
          "format": "string",
          "description": "The direction in which the security group rule is applied. The only values\nallowed are \"ingress\" or \"egress\". For a compute instance, an ingress\nsecurity group rule is applied to incoming (ingress) traffic for that\ninstance. An egress rule is applied to traffic leaving the instance."
        },
        "ethertype": {
          "type": "string",
          "format": "string",
          "title": "Must be IPv4 or IPv6, and addresses represented in CIDR must match the\ningress or egress rules.\nEtherType string json:\"ethertype\""
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "The UUID for this security group rule."
        },
        "port_range_max": {
          "type": "integer",
          "format": "int32",
          "description": "The maximum port number in the range that is matched by the security group\nrule. The PortRangeMin attribute constrains the PortRangeMax attribute. If\nthe protocol is ICMP, this value must be an ICMP type."
        },
        "port_range_min": {
          "type": "integer",
          "format": "int32",
          "description": "The minimum port number in the range that is matched by the security group\nrule. If the protocol is TCP or UDP, this value must be less than or equal\nto the value of the PortRangeMax attribute. If the protocol is ICMP, this\nvalue must be an ICMP type."
        },
        "protocol": {
          "type": "string",
          "format": "string",
          "description": "The protocol that is matched by the security group rule. Valid values are\n\"tcp\", \"udp\", \"icmp\" or an empty string."
        },
        "remote_group_id": {
          "type": "string",
          "format": "string",
          "description": "The remote group ID to be associated with this security group rule. You\ncan specify either RemoteGroupID or RemoteIPPrefix."
        },
        "remote_ip_prefix": {
          "type": "string",
          "format": "string",
          "description": "The remote IP prefix to be associated with this security group rule. You\ncan specify either RemoteGroupID or RemoteIPPrefix . This attribute\nmatches the specified IP prefix as the source IP address of the IP packet."
        },
        "security_group_id": {
          "type": "string",
          "format": "string",
          "description": "The security group ID to associate with this security group rule."
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "description": "The owner of this security group rule."
        }
      },
      "description": "SecGroupRule represents a rule to dictate the behaviour of incoming or\noutgoing traffic for a particular security group."
    },
    "neutronSubnet": {
      "type": "object",
      "properties": {
        "allocation_pools": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronAllocationPool"
          }
        },
        "cidr": {
          "type": "string",
          "format": "string"
        },
        "dns_name_servers": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "enable_dhcp": {
          "type": "boolean",
          "format": "boolean"
        },
        "gateway_ip": {
          "type": "string",
          "format": "string"
        },
        "host_routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronHostRoute"
          }
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "ip_version": {
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "network_id": {
          "type": "string",
          "format": "string"
        },
        "tenant_id": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "novaAddress": {
      "type": "object",
      "properties": {
        "addr": {
          "type": "string",
          "format": "string"
        },
        "assigned_type": {
          "type": "string",
          "format": "string"
        },
        "ext": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "byte"
          }
        },
        "mac_addr": {
          "type": "string",
          "format": "string"
        },
        "version": {
          "type": "integer",
          "format": "int32"
        }
      },
      "description": "Address represents an IP address."
    },
    "novaAddresses": {
      "type": "object",
      "properties": {
        "addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/novaAddress"
          }
        }
      }
    },
    "novaFile": {
      "type": "object",
      "properties": {
        "contents": {
          "type": "string",
          "format": "byte",
          "description": "Contents of the file. Maximum content size is 255 bytes."
        },
        "path": {
          "type": "string",
          "format": "string",
          "title": "Path of the file"
        }
      },
      "description": "File is used within CreateOpts and RebuildOpts to inject a file into the server at launch.\nFile implements the json.Marshaler interface, so when a Create or Rebuild operation is requested,\njson.Marshal will call File's MarshalJSON method."
    },
    "novaGroup": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "tenant_id": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "Group represents a group."
    },
    "novaIPRange": {
      "type": "object",
      "properties": {
        "cidr": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "IPRange represents the IP range whose traffic will be accepted by the\nsecurity group."
    },
    "novaRule": {
      "type": "object",
      "properties": {
        "from_port": {
          "type": "integer",
          "format": "int32",
          "title": "The lower bound of the port range which this security group should open up"
        },
        "group": {
          "$ref": "#/definitions/novaGroup",
          "description": "Not documented."
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "The unique ID. If Neutron is installed, this ID will be\nrepresented as a string UUID; if Neutron is not installed, it will be a\nnumeric ID. For the sake of consistency, we always cast it to a string."
        },
        "ip_protocol": {
          "type": "string",
          "format": "string",
          "title": "The IP protocol (e.g. TCP) which the security group accepts"
        },
        "ip_range": {
          "$ref": "#/definitions/novaIPRange",
          "title": "The CIDR IP range whose traffic can be received"
        },
        "parent_group_id": {
          "type": "string",
          "format": "string",
          "title": "The security group ID to which this rule belongs"
        },
        "to_port": {
          "type": "integer",
          "format": "int32",
          "title": "The upper bound of the port range which this security group should open up"
        }
      },
      "description": "Rule represents a security group rule, a policy which determines how a\nsecurity group operates and what inbound traffic it allows in."
    },
    "novaSecurityGroup": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "format": "string",
          "description": "The human-readable description of the group."
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "The unique ID of the group. If Neutron is installed, this ID will be\nrepresented as a string UUID; if Neutron is not installed, it will be a\nnumeric ID. For the sake of consistency, we always cast it to a string."
        },
        "name": {
          "type": "string",
          "format": "string",
          "description": "The human-readable name of the group, which needs to be unique."
        },
        "rules": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/novaRule"
          },
          "description": "The rules which determine how this security group operates."
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "description": "The ID of the tenant to which this security group belongs."
        }
      },
      "description": "SecurityGroup represents a security group."
    },
    "novaSecurityGroups": {
      "type": "object",
      "properties": {
        "security_groups": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/novaSecurityGroup"
          }
        }
      }
    },
    "novaServer": {
      "type": "object",
      "properties": {
        "accessIPv4": {
          "type": "string",
          "format": "string",
          "description": "AccessIPv4 and AccessIPv6 contain the IP addresses of the server, suitable for remote access for administration."
        },
        "accessIPv6": {
          "type": "string",
          "format": "string"
        },
        "addresses": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/novaAddresses"
          },
          "description": "Addresses includes a list of all IP addresses assigned to the server, keyed by pool."
        },
        "adminPass": {
          "type": "string",
          "format": "string",
          "description": "AdminPass will generally be empty (\"\").  However, it will contain the administrative password chosen when provisioning a new server without a set AdminPass setting in the first place.\nNote that this is the ONLY time this field will be valid."
        },
        "created": {
          "type": "string",
          "format": "string",
          "title": "google.protobuf.Timestamp created = 6;"
        },
        "flavors": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/openstacknovaFlavor"
          },
          "description": "Flavor refers to a JSON object, which itself indicates the hardware configuration of the deployed server."
        },
        "host_id": {
          "type": "string",
          "format": "string"
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "ID uniquely identifies this server amongst all other servers, including those not accessible to the current tenant."
        },
        "images": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/openstacknovaImage"
          },
          "description": "Image refers to a JSON object, which itself indicates the OS image used to deploy the server."
        },
        "key_name": {
          "type": "string",
          "format": "string",
          "description": "KeyName indicates which public key was injected into the server on launch."
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "description": "Links includes HTTP references to the itself, useful for passing along to other APIs that might want a server reference."
        },
        "metadata_info": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "description": "Metadata includes a list of all user-specified key-value pairs attached to the server."
        },
        "name": {
          "type": "string",
          "format": "string",
          "description": "Name contains the human-readable name for the server."
        },
        "progress": {
          "type": "integer",
          "format": "int32",
          "description": "Progress ranges from 0..100.\nA request made against the server completes only once Progress reaches 100."
        },
        "security_groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/novaSecurityGroups"
          },
          "title": "SecurityGroups includes the security groups that this instance has applied to it"
        },
        "status": {
          "type": "string",
          "format": "string",
          "description": "Status contains the current operational status of the server, such as IN_PROGRESS or ACTIVE."
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "description": "TenantID identifies the tenant owning this server resource."
        },
        "updated": {
          "type": "string",
          "format": "string",
          "title": "Updated and Created contain ISO-8601 timestamps of when the state of the server last changed, and when it was created.\ngoogle.protobuf.Timestamp updated = 5;"
        },
        "user_id": {
          "type": "string",
          "format": "string",
          "description": "UserID uniquely identifies the user account owning the tenant."
        }
      },
      "description": "Server exposes only the standard OpenStack fields corresponding to a given server on the user's account."
    },
    "openstackConsoleResourceRequestData": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackConsoleResourceResponseData": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackFlavorDiscoveryReqRespData": {
      "type": "object",
      "properties": {
        "access_type": {
          "type": "string",
          "format": "string"
        },
        "flavors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstacknovaFlavor"
          }
        }
      },
      "description": "access_type: means 'public', ...\nflavors: return value, all of flavors information",
      "title": "Flavor data structure, as input/output argument"
    },
    "openstackIdNamePair": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackIfaceInfo": {
      "type": "object",
      "properties": {
        "interface_info_id": {
          "type": "string",
          "format": "string"
        },
        "network_id": {
          "type": "string",
          "format": "string"
        },
        "network_name": {
          "type": "string",
          "format": "string"
        },
        "port_id": {
          "type": "string",
          "format": "string"
        },
        "port_name": {
          "type": "string",
          "format": "string"
        },
        "router_id": {
          "type": "string",
          "format": "string"
        },
        "router_name": {
          "type": "string",
          "format": "string"
        },
        "secgroups_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackSecGroupInfo"
          }
        },
        "subnet_id": {
          "type": "string",
          "format": "string"
        },
        "subnet_name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackImageDiscoveryReqRespData": {
      "type": "object",
      "properties": {
        "images": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/echopbopenstackImage"
          }
        },
        "member_status": {
          "type": "string",
          "format": "string"
        },
        "status": {
          "type": "string",
          "format": "string"
        },
        "visibility": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "status: means 'queued', 'saving', 'active', ...\nvisibility: means 'public', 'shared', 'community', ...\nmember_status: means 'accepted', 'all', ...\nimages: return value, all of images Information",
      "title": "Image data structure, as input/output argument"
    },
    "openstackLibvirtDomainInfo": {
      "type": "object",
      "properties": {
        "display": {
          "type": "string",
          "format": "string"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "uuid": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "For Libvirt"
    },
    "openstackLibvirtDomainReqRespData": {
      "type": "object",
      "properties": {
        "domain_info": {
          "$ref": "#/definitions/openstackLibvirtDomainInfo"
        },
        "server_id": {
          "type": "string",
          "format": "string"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackMachineDestroyReqRespData": {
      "type": "object",
      "properties": {
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        },
        "vms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackIdNamePair"
          }
        }
      },
      "description": "id: required if name is empty\nname: required if id is empty\nstate_: return value",
      "title": "Machine data structure, as input/output argument"
    },
    "openstackMachineDiscoveryReqRespData": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "format": "string"
        },
        "vms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/novaServer"
          }
        }
      },
      "description": "status: means 'ACTIVE', ...\nflavors: return value, all of flavors information",
      "title": "Machine data structure, as input/output argument"
    },
    "openstackMachineRebootReqRespData": {
      "type": "object",
      "properties": {
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        },
        "vms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackIdNamePair"
          }
        }
      },
      "description": "id: required if name is empty\nname: required if id is empty\nstate_: return value",
      "title": "Machine data structure, as input/output argument"
    },
    "openstackMachineSpawnsReqRespData": {
      "type": "object",
      "properties": {
        "vms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackOpenstackNovaBootReqRespData"
          }
        }
      },
      "description": "OpenstackNovaBootReqRespData: see definition of \"OpenstackNovaBootReqRespData\"",
      "title": "Machine data structure, as input/output argument"
    },
    "openstackNetworkDiscoveryReqRespData": {
      "type": "object",
      "properties": {
        "networks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackneutronNetwork"
          }
        },
        "status": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "status: means 'ACTIVE', ...\nnetworks: return value, all of networks information",
      "title": "Network data structure, as input/output argument"
    },
    "openstackNetworkTopologyReqRespData": {
      "type": "object",
      "properties": {
        "floating_network_id": {
          "type": "string",
          "format": "string"
        },
        "information": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackOpenstackNeutronLandscapeReqRespData"
          }
        }
      },
      "description": "networks: argument, list of network (id or name)\nnetworks: return value, all of networks information",
      "title": "Network data structure, as input/output argument"
    },
    "openstackOpenstackNeutronLandscapeReqRespData": {
      "type": "object",
      "properties": {
        "gateways_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/echopbopenstackGatewayInfo"
          }
        },
        "ifaces_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackIfaceInfo"
          }
        },
        "interfaces_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronInterfaceInfo"
          }
        },
        "ports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronPort"
          }
        },
        "secgroup": {
          "$ref": "#/definitions/neutronSecGroup"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        },
        "vnets": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackneutronNetwork"
          }
        },
        "vrouter": {
          "$ref": "#/definitions/neutronRouter"
        }
      }
    },
    "openstackOpenstackNeutronNetRequestData": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "network": {
          "$ref": "#/definitions/openstackneutronNetwork"
        },
        "router": {
          "$ref": "#/definitions/neutronRouter"
        }
      }
    },
    "openstackOpenstackNeutronNetResponseData": {
      "type": "object",
      "properties": {
        "gateway_network_id": {
          "type": "string",
          "format": "string"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "interface_info": {
          "$ref": "#/definitions/neutronInterfaceInfo"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "network": {
          "$ref": "#/definitions/openstackneutronNetwork"
        },
        "router": {
          "$ref": "#/definitions/neutronRouter"
        }
      }
    },
    "openstackOpenstackNovaBootReqRespData": {
      "type": "object",
      "properties": {
        "fip_server_pairs": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "FloatingIP-Server dictionary\nReturn value"
        },
        "flavor_id": {
          "type": "string",
          "format": "string",
          "title": "Flavor id, command line 'openstack flavor list' or 'nova flavor-list' to get alls\nRequired if 'flavor_name' is not provided"
        },
        "flavor_name": {
          "type": "string",
          "format": "string",
          "title": "Flavor name\nRequired if 'flavor_id' is not provided"
        },
        "floating_ips": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronFloatingIP"
          },
          "title": "Floating IP details\nReturn value"
        },
        "floating_network_id": {
          "type": "string",
          "format": "string",
          "title": "The external network that floating ip should generate into\nRequired, if 'floating_network_name' is not provided"
        },
        "floating_network_name": {
          "type": "string",
          "format": "string",
          "title": "The internal network that machine should boot into\nRequired alternatively for 'floating_network_id'"
        },
        "image_id": {
          "type": "string",
          "format": "string",
          "title": "Image id, command line 'openstack image list' or 'glance image-list' to get alls\nRequired if 'image_name' is not provided"
        },
        "image_name": {
          "type": "string",
          "format": "string",
          "title": "Image name\nRequired if 'image_id' is not provided"
        },
        "max_count": {
          "type": "integer",
          "format": "int32",
          "title": "Max count, the maxmized machines to boot\nOptional, the default just one machine, must greater or equal than 'min_count'"
        },
        "min_count": {
          "type": "integer",
          "format": "int32",
          "title": "Min count, the minimized machines to boot\nOptional, the default just one machine, must less or equal than 'max_count'"
        },
        "name_prefix": {
          "type": "string",
          "format": "string",
          "title": "Server name, or name prefix for multi-creations\nRequired"
        },
        "network_id": {
          "type": "string",
          "format": "string",
          "title": "The internal network that machine should boot into\nRequired, if 'network_name' is not provided"
        },
        "network_name": {
          "type": "string",
          "format": "string",
          "title": "The internal network that machine should boot into\nRequired alternatively for 'network_id'"
        },
        "personality": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/novaFile"
          },
          "title": "The upload file (bytes) inject into booting\nOptional"
        },
        "port_server_pairs": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Port-Server dictionary\nReturn value"
        },
        "ports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronPort"
          },
          "title": "Port details, the nic content created for server\nReturn value"
        },
        "secgroups_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openstackSecGroupInfo"
          },
          "title": "Security Groups Information, array of ID/Name pairs\nOptional, nova will using 'default' secgroup"
        },
        "servers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/novaServer"
          },
          "title": "Server details\nReturn value"
        },
        "state_code": {
          "type": "integer",
          "format": "int32",
          "title": "State code, return non zero decimal if creation failed\nReturn value"
        },
        "state_message": {
          "type": "string",
          "format": "string",
          "title": "State message, return errors if createion failed"
        },
        "user_data": {
          "type": "string",
          "format": "byte",
          "title": "User data to inject into booting\nOptional"
        }
      },
      "description": "Booting machines data structure, used as input/output argument.\n\nFor request, only 'flavor', 'image', 'count', 'secgroups', 'networks' are required\nAs response, provide 'servers' with additonal 'port', 'floatingip' relations"
    },
    "openstackSecGroupInfo": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackSubnetDiscoveryReqRespData": {
      "type": "object",
      "properties": {
        "network_id": {
          "type": "string",
          "format": "string"
        },
        "subnets": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronSubnet"
          }
        }
      },
      "description": "network_id: search condition\nsubnets: return value, all of subnets information",
      "title": "Subnet data structure, as input/output argument"
    },
    "openstackTokenReqRespData": {
      "type": "object",
      "properties": {
        "src": {
          "$ref": "#/definitions/identityTokenCredentialsV2"
        },
        "tgt": {
          "$ref": "#/definitions/identityToken"
        }
      }
    },
    "openstackneutronGatewayInfo": {
      "type": "object",
      "properties": {
        "network_id": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "openstackneutronNetwork": {
      "type": "object",
      "properties": {
        "admin_state_up": {
          "type": "boolean",
          "format": "boolean",
          "title": "The administrative state of network. If false (down), the network does not forward packets.\nAdminStateUp bool json:\"admin_state_up\""
        },
        "id": {
          "type": "string",
          "format": "string",
          "title": "UUID for the network\nID string json:\"id\""
        },
        "name": {
          "type": "string",
          "format": "string",
          "title": "Human-readable name for the network. Might not be unique.\nName string json:\"name\""
        },
        "shared": {
          "type": "boolean",
          "format": "boolean",
          "title": "Specifies whether the network resource can be accessed by any tenant or not.\nShared bool json:\"shared\""
        },
        "status": {
          "type": "string",
          "format": "string",
          "title": "Indicates whether network is currently operational. Possible values include\n'ACTIVE', 'DOWN', 'BUILD', or 'ERROR'. Plug-ins might define additional values.\nStatus string json:\"status\""
        },
        "subnets": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/neutronSubnet"
          },
          "title": "Subnets associated with this network.\nSubnets []string json:\"subnets\""
        },
        "tenant_id": {
          "type": "string",
          "format": "string",
          "title": "Owner of network. Only admin users can specify a tenant_id other than its own.\nTenantID string json:\"tenant_id\""
        }
      },
      "description": "Network represents, well, a network."
    },
    "openstacknovaFlavor": {
      "type": "object",
      "properties": {
        "disk": {
          "type": "integer",
          "format": "int32",
          "title": "The Disk and RA\u003c fields provide a measure of storage space offered by the flavor, in GB and MB, respectively.\nDisk int json:\"disk\"\nRAM  int json:\"ram\""
        },
        "id": {
          "type": "string",
          "format": "string",
          "title": "The Id field contains the flavor's unique identifier.\nFor example, this identifier will be useful when specifying which hardware configuration to use for a new server instance.\nID string json:\"id\""
        },
        "is_public": {
          "type": "boolean",
          "format": "boolean",
          "title": "IsPublic indicates whether the flavor is public.\nIsPublic bool json:\"is_public\""
        },
        "name": {
          "type": "string",
          "format": "string",
          "title": "The Name field provides a human-readable moniker for the flavor.\nName       string  json:\"name\"\nRxTxFactor float64 json:\"rxtx_factor\""
        },
        "ram": {
          "type": "integer",
          "format": "int32"
        },
        "rxtx_factor": {
          "type": "number",
          "format": "double"
        },
        "swap": {
          "type": "integer",
          "format": "int32",
          "title": "Swap indicates how much space is reserved for swap.\nIf not provided, this field will be set to 0.\nSwap int json:\"swap\""
        },
        "vcpus": {
          "type": "integer",
          "format": "int32",
          "title": "VCPUs indicates how many (virtual) CPUs are available for this flavor.\nVCPUs int json:\"vcpus\""
        }
      },
      "description": "Flavor records represent (virtual) hardware configurations for server resources in a region."
    },
    "openstacknovaImage": {
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "string"
        },
        "id": {
          "type": "string",
          "format": "string",
          "description": "ID contains the image's unique identifier."
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "min_disk": {
          "type": "integer",
          "format": "int32",
          "description": "MinDisk and MinRAM specify the minimum resources a server must provide to be able to install the image."
        },
        "min_ram": {
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string",
          "format": "string",
          "description": "Name provides a human-readable moniker for the OS image."
        },
        "progress": {
          "type": "integer",
          "format": "int32",
          "description": "The Progress and Status fields indicate image-creation status.\nAny usable image will have 100% progress."
        },
        "status": {
          "type": "string",
          "format": "string"
        },
        "updated": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "Image is used for JSON (un)marshalling.\nIt provides a description of an OS image."
    }
  }
}
`
)
