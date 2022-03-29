package schema

// AlexaSmartHome is the schema for smart home requests/responses
const AlexaSmartHome = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "Alexa Smart Home Message Schema",
    "description": "A JSON message sent from a skill to Alexa, either proactively or as a response to a directive",
    "definitions": {
        "common.properties": {
            "payloadVersion": {
                "enum": [
                    "3"
                ]
            },
            "message": {
                "type": "string"
            },
            "currentDeviceMode": {
                "enum": [
                    "ASLEEP",
                    "NOT_PROVISIONED",
                    "COLOR",
                    "OTHER"
                ]
            },
            "messageId": {
                "type": "string",
                "minLength": 1,
                "maxLength": 127,
                "pattern": "^[a-zA-Z0-9\\-]*$"
            },
            "correlationToken": {
                "type": "string",
                "minLength": 1
            },
            "temperature": {
                "type": "object",
                "required": [
                    "value",
                    "scale"
                ],
                "additionalProperties": false,
                "properties": {
                    "value": {
                        "type": "number"
                    },
                    "scale": {
                        "enum": [
                            "CELSIUS",
                            "FAHRENHEIT",
                            "KELVIN"
                        ]
                    }
                }
            },
            "channel": {
                "type": "object",
                "additionalProperties": false,
                "properties": {
                    "number": {
                        "type": "string"
                    },
                    "callSign": {
                        "type": "string"
                    },
                    "affiliateCallSign": {
                        "type": "string"
                    }
                },
                "minProperties": 1
            },
            "input": {
                "type": "string",
                "minLength": 1
            },
            "volume": {
                "type": "integer",
                "minimum": 0,
                "maximum": 100
            },
            "muted": {
                "type": "boolean"
            },
            "timestamp": {
                "type": "string",
                "format": "date-time"
            },
            "uncertaintyInMilliseconds": {
                "type": "integer",
                "minimum": 0
            },
            "resolution": {
                "type": "object",
                "required": [
                    "width",
                    "height"
                ],
                "additionalProperties": false,
                "properties": {
                    "width": {
                        "type": "integer",
                        "minimum": 0
                    },
                    "height": {
                        "type": "integer",
                        "minimum": 0
                    }
                }
            },
            "protocol": {
                "enum": [
                    "RTSP"
                ]
            },
            "authorizationType": {
                "enum": [
                    "BASIC",
                    "DIGEST",
                    "NONE"
                ]
            },
            "videoCodec": {
                "enum": [
                    "H264",
                    "MPEG2",
                    "MJPEG",
                    "JPG"
                ]
            },
            "audioCodec": {
                "enum": [
                    "G711",
                    "AAC",
                    "NONE"
                ]
            },
            "cameraStream": {
                "type": "object",
                "required": [
                    "uri",
                    "protocol",
                    "resolution",
                    "authorizationType",
                    "videoCodec",
                    "audioCodec"
                ],
                "additionalProperties": false,
                "properties": {
                    "uri": {
                        "type": "string",
                        "format": "uri"
                    },
                    "expirationTime": {
                        "type": "string",
                        "format": "date-time"
                    },
                    "idleTimeoutSeconds": {
                        "type": "integer",
                        "minimum": 0
                    },
                    "protocol": {
                        "$ref": "#/definitions/common.properties/protocol"
                    },
                    "resolution": {
                        "$ref": "#/definitions/common.properties/resolution"
                    },
                    "authorizationType": {
                        "$ref": "#/definitions/common.properties/authorizationType"
                    },
                    "videoCodec": {
                        "$ref": "#/definitions/common.properties/videoCodec"
                    },
                    "audioCodec": {
                        "$ref": "#/definitions/common.properties/audioCodec"
                    }
                }
            },
            "cause": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "APP_INTERACTION",
                            "PHYSICAL_INTERACTION",
                            "PERIODIC_POLL",
                            "RULE_TRIGGER",
                            "VOICE_INTERACTION"
                        ]
                    }
                }
            },
            "version": {
                "enum": [
                    "3"
                ]
            },
            "endpointId": {
                "type": "string",
                "pattern": "^[a-zA-Z0-9_\\-=#;:?@&]*$",
                "minLength": 1,
                "maxLength": 256
            },
            "endpoint": {
                "type": "object",
                "required": [
                    "endpointId"
                ],
                "properties": {
                    "scope": {
                        "type": "object",
                        "required": [
                            "type",
                            "token"
                        ],
                        "properties": {
                            "type": {
                                "enum": [
                                    "BearerToken"
                                ]
                            },
                            "token": {
                                "type": "string",
                                "minLength": 1
                            }
                        }
                    },
                    "endpointId": {
                        "$ref": "#/definitions/common.properties/endpointId"
                    }
                }
            },
            "interfaces": {
                "Alexa": {
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            }
                        }
                    }
                },
                "EndpointHealth": {
                    "connectivity": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.EndpointHealth"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "connectivity"
                                    ]
                                },
                                "value": {
                                    "type": "object",
                                    "required": [
                                        "value"
                                    ],
                                    "additionalProperties": false,
                                    "properties": {
                                        "value": {
                                            "enum": [
                                                "OK",
                                                "UNREACHABLE"
                                            ]
                                        }
                                    }
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.EndpointHealth"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "connectivity"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "PowerController": {
                    "powerState": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.PowerController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "powerState"
                                    ]
                                },
                                "value": {
                                    "enum": [
                                        "ON",
                                        "OFF"
                                    ]
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.PowerController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "powerState"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "PowerLevelController": {
                    "powerLevel": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.PowerLevelController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "powerLevel"
                                    ]
                                },
                                "value": {
                                    "type": "integer",
                                    "minimum": 0,
                                    "maximum": 100
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.PowerLevelController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "powerLevel"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "PercentageController": {
                    "percentage": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.PercentageController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "percentage"
                                    ]
                                },
                                "value": {
                                    "type": "integer",
                                    "minimum": 0,
                                    "maximum": 100
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.PercentageController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "percentage"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "BrightnessController": {
                    "brightness": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.BrightnessController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "brightness"
                                    ]
                                },
                                "value": {
                                    "type": "integer",
                                    "minimum": 0,
                                    "maximum": 100
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.BrightnessController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "brightness"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "CameraStreamController": {
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version",
                            "cameraStreamConfigurations"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.CameraStreamController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "cameraStreamConfigurations": {
                                "type": "array",
                                "uniqueItems": true,
                                "items": {
                                    "type": "object",
                                    "required": [
                                        "protocols",
                                        "resolutions",
                                        "authorizationTypes",
                                        "videoCodecs",
                                        "audioCodecs"
                                    ],
                                    "additionalProperties": false,
                                    "properties": {
                                        "protocols": {
                                            "type": "array",
                                            "uniqueItems": true,
                                            "minItems": 1,
                                            "items": {
                                                "$ref": "#/definitions/common.properties/protocol"
                                            }
                                        },
                                        "resolutions": {
                                            "type": "array",
                                            "uniqueItems": true,
                                            "minItems": 1,
                                            "items": {
                                                "$ref": "#/definitions/common.properties/resolution"
                                            }
                                        },
                                        "authorizationTypes": {
                                            "type": "array",
                                            "uniqueItems": true,
                                            "minItems": 1,
                                            "items": {
                                                "$ref": "#/definitions/common.properties/authorizationType"
                                            }
                                        },
                                        "videoCodecs": {
                                            "type": "array",
                                            "uniqueItems": true,
                                            "minItems": 1,
                                            "items": {
                                                "$ref": "#/definitions/common.properties/videoCodec"
                                            }
                                        },
                                        "audioCodecs": {
                                            "type": "array",
                                            "uniqueItems": true,
                                            "minItems": 1,
                                            "items": {
                                                "$ref": "#/definitions/common.properties/audioCodec"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                "ColorController": {
                    "color": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.ColorController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "color"
                                    ]
                                },
                                "value": {
                                    "type": "object",
                                    "required": [
                                        "hue",
                                        "saturation",
                                        "brightness"
                                    ],
                                    "additionalProperties": false,
                                    "properties": {
                                        "hue": {
                                            "type": "number",
                                            "minimum": 0,
                                            "maximum": 360
                                        },
                                        "saturation": {
                                            "type": "number",
                                            "minimum": 0,
                                            "maximum": 1
                                        },
                                        "brightness": {
                                            "type": "number",
                                            "minimum": 0,
                                            "maximum": 1
                                        }
                                    }
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.ColorController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "color"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "ColorTemperatureController": {
                    "colorTemperatureInKelvin": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.ColorTemperatureController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "colorTemperatureInKelvin"
                                    ]
                                },
                                "value": {
                                    "type": "integer",
                                    "minimum": 1000,
                                    "maximum": 10000
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.ColorTemperatureController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "colorTemperatureInKelvin"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "LockController": {
                    "lockState": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.LockController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "lockState"
                                    ]
                                },
                                "value": {
                                    "enum": [
                                        "LOCKED",
                                        "UNLOCKED",
                                        "JAMMED"
                                    ]
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.LockController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "lockState"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "ThermostatController": {
                    "setpoint": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.ThermostatController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "targetSetpoint",
                                        "lowerSetpoint",
                                        "upperSetpoint"
                                    ]
                                },
                                "value": {
                                    "$ref": "#/definitions/common.properties/temperature"
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "thermostatMode": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.ThermostatController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "thermostatMode"
                                    ]
                                },
                                "value": {
                                    "enum": [
                                        "HEAT",
                                        "COOL",
                                        "AUTO",
                                        "ECO",
                                        "OFF",
                                        "CUSTOM"
                                    ]
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.ThermostatController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "targetSetpoint",
                                                        "lowerSetpoint",
                                                        "upperSetpoint",
                                                        "thermostatMode"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "TemperatureSensor": {
                    "temperature": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.TemperatureSensor"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "temperature"
                                    ]
                                },
                                "value": {
                                    "$ref": "#/definitions/common.properties/temperature"
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.TemperatureSensor"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "temperature"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "ChannelController": {
                    "channel": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.ChannelController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "channel"
                                    ]
                                },
                                "value": {
                                    "$ref": "#/definitions/common.properties/channel"
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.ChannelController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "channel"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "InputController": {
                    "input": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.InputController"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "input"
                                    ]
                                },
                                "value": {
                                    "$ref": "#/definitions/common.properties/input"
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.InputController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "input"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "Speaker": {
                    "volume": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.Speaker"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "volume"
                                    ]
                                },
                                "value": {
                                    "$ref": "#/definitions/common.properties/volume"
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "muted": {
                        "property": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "value",
                                "timeOfSample",
                                "uncertaintyInMilliseconds"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.Speaker"
                                    ]
                                },
                                "name": {
                                    "enum": [
                                        "muted"
                                    ]
                                },
                                "value": {
                                    "$ref": "#/definitions/common.properties/muted"
                                },
                                "timeOfSample": {
                                    "$ref": "#/definitions/common.properties/timestamp"
                                },
                                "uncertaintyInMilliseconds": {
                                    "$ref": "#/definitions/common.properties/uncertaintyInMilliseconds"
                                }
                            }
                        }
                    },
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.Speaker"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "properties": {
                                "type": "object",
                                "required": [
                                    "supported",
                                    "proactivelyReported",
                                    "retrievable"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "supported": {
                                        "type": "array",
                                        "uniqueItems": true,
                                        "items": {
                                            "type": "object",
                                            "required": [
                                                "name"
                                            ],
                                            "additionalProperties": false,
                                            "properties": {
                                                "name": {
                                                    "enum": [
                                                        "volume",
                                                        "muted"
                                                    ]
                                                }
                                            }
                                        }
                                    },
                                    "proactivelyReported": {
                                        "type": "boolean"
                                    },
                                    "retrievable": {
                                        "type": "boolean"
                                    }
                                }
                            }
                        }
                    }
                },
                "SceneController": {
                    "capabilities": {
                        "type": "object",
                        "required": [
                            "type",
                            "interface",
                            "version",
                            "supportsDeactivation",
                            "proactivelyReported"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "type": {
                                "enum": [
                                    "AlexaInterface"
                                ]
                            },
                            "interface": {
                                "enum": [
                                    "Alexa.SceneController"
                                ]
                            },
                            "version": {
                                "$ref": "#/definitions/common.properties/version"
                            },
                            "supportsDeactivation": {
                                "type": "boolean"
                            },
                            "proactivelyReported": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "state.properties": {
                "type": "array",
                "uniqueItems": true,
                "additionalItems": false,
                "items": {
                    "anyOf": [
                        {
                            "$ref": "#/definitions/common.properties/interfaces/EndpointHealth/connectivity/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/PowerController/powerState/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/PowerLevelController/powerLevel/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/PercentageController/percentage/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/BrightnessController/brightness/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/ColorController/color/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/ColorTemperatureController/colorTemperatureInKelvin/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/LockController/lockState/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/ThermostatController/setpoint/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/ThermostatController/thermostatMode/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/TemperatureSensor/temperature/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/ChannelController/channel/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/InputController/input/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/Speaker/volume/property"
                        },
                        {
                            "$ref": "#/definitions/common.properties/interfaces/Speaker/muted/property"
                        }
                    ]
                }
            },
            "context": {
                "type": "object",
                "additionalProperties": false,
                "properties": {
                    "properties": {
                        "$ref": "#/definitions/common.properties/state.properties"
                    }
                }
            },
            "payload": {
                "cameraStreams": {
                    "type": "object",
                    "additionalProperties": false,
                    "properties": {
                        "cameraStreams": {
                            "type": "array",
                            "uniqueItems": true,
                            "minItems": 1,
                            "items": {
                                "$ref": "#/definitions/common.properties/cameraStream"
                            }
                        },
                        "imageUri": {
                            "type": "string",
                            "format": "uri"
                        }
                    }
                },
                "sceneActivationDeactivation": {
                    "type": "object",
                    "required": [
                        "cause",
                        "timestamp"
                    ],
                    "additionalProperties": false,
                    "properties": {
                        "cause": {
                            "$ref": "#/definitions/common.properties/cause"
                        },
                        "timestamp": {
                            "$ref": "#/definitions/common.properties/timestamp"
                        }
                    }
                }
            }
        },
        "ErrorResponse.properties": {
            "name": {
                "enum": [
                    "ErrorResponse"
                ]
            },
            "header.general": {
                "type": "object",
                "required": [
                    "namespace",
                    "name",
                    "payloadVersion",
                    "messageId"
                ],
                "additionalProperties": false,
                "properties": {
                    "namespace": {
                        "enum": [
                            "Alexa"
                        ]
                    },
                    "name": {
                        "$ref": "#/definitions/ErrorResponse.properties/name"
                    },
                    "payloadVersion": {
                        "$ref": "#/definitions/common.properties/payloadVersion"
                    },
                    "messageId": {
                        "$ref": "#/definitions/common.properties/messageId"
                    },
                    "correlationToken": {
                        "$ref": "#/definitions/common.properties/correlationToken"
                    }
                }
            },
            "header.ThermostatController": {
                "type": "object",
                "required": [
                    "namespace",
                    "name",
                    "payloadVersion",
                    "messageId"
                ],
                "additionalProperties": false,
                "properties": {
                    "namespace": {
                        "enum": [
                            "Alexa.ThermostatController"
                        ]
                    },
                    "name": {
                        "$ref": "#/definitions/ErrorResponse.properties/name"
                    },
                    "payloadVersion": {
                        "$ref": "#/definitions/common.properties/payloadVersion"
                    },
                    "messageId": {
                        "$ref": "#/definitions/common.properties/messageId"
                    },
                    "correlationToken": {
                        "$ref": "#/definitions/common.properties/correlationToken"
                    }
                }
            },
            "header.ColorTemperatureController": {
                "type": "object",
                "required": [
                    "namespace",
                    "name",
                    "payloadVersion",
                    "messageId"
                ],
                "additionalProperties": false,
                "properties": {
                    "namespace": {
                        "enum": [
                            "Alexa.ColorTemperatureController"
                        ]
                    },
                    "name": {
                        "$ref": "#/definitions/ErrorResponse.properties/name"
                    },
                    "payloadVersion": {
                        "$ref": "#/definitions/common.properties/payloadVersion"
                    },
                    "messageId": {
                        "$ref": "#/definitions/common.properties/messageId"
                    },
                    "correlationToken": {
                        "$ref": "#/definitions/common.properties/correlationToken"
                    }
                }
            },
            "header.Authorization": {
                "type": "object",
                "required": [
                    "namespace",
                    "name",
                    "payloadVersion",
                    "messageId"
                ],
                "additionalProperties": false,
                "properties": {
                    "namespace": {
                        "enum": [
                            "Alexa.Authorization"
                        ]
                    },
                    "name": {
                        "$ref": "#/definitions/ErrorResponse.properties/name"
                    },
                    "payloadVersion": {
                        "$ref": "#/definitions/common.properties/payloadVersion"
                    },
                    "messageId": {
                        "$ref": "#/definitions/common.properties/messageId"
                    },
                    "correlationToken": {
                        "$ref": "#/definitions/common.properties/correlationToken"
                    }
                }
            },
            "payload.general": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "BRIDGE_UNREACHABLE",
                            "ENDPOINT_BUSY",
                            "ENDPOINT_LOW_POWER",
                            "ENDPOINT_UNREACHABLE",
                            "EXPIRED_AUTHORIZATION_CREDENTIAL",
                            "FIRMWARE_OUT_OF_DATE",
                            "HARDWARE_MALFUNCTION",
                            "INTERNAL_ERROR",
                            "INVALID_AUTHORIZATION_CREDENTIAL",
                            "INVALID_DIRECTIVE",
                            "INVALID_VALUE",
                            "NO_SUCH_ENDPOINT",
                            "RATE_LIMIT_EXCEEDED",
                            "ACCEPT_GRANT_FAILED"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    }
                }
            },
            "payload.general.VALUE_OUT_OF_RANGE": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "VALUE_OUT_OF_RANGE"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    },
                    "validRange": {
                        "type": "object",
                        "additionalProperties": false,
                        "properties": {
                            "minimumValue": {
                                "type": "number"
                            },
                            "maximumValue": {
                                "type": "number"
                            }
                        }
                    }
                }
            },
            "payload.general.NOT_SUPPORTED_IN_CURRENT_MODE": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "NOT_SUPPORTED_IN_CURRENT_MODE"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    },
                    "currentDeviceMode": {
                        "$ref": "#/definitions/common.properties/currentDeviceMode"
                    }
                }
            },
            "payload.general.ENDPOINT_LOW_POWER": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "ENDPOINT_LOW_POWER"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    }
                }
            },
            "payload.general.TEMPERATURE_VALUE_OUT_OF_RANGE": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "TEMPERATURE_VALUE_OUT_OF_RANGE"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    },
                    "validRange": {
                        "type": "object",
                        "additionalProperties": false,
                        "properties": {
                            "minimumValue": {
                                "$ref": "#/definitions/common.properties/temperature"
                            },
                            "maximumValue": {
                                "$ref": "#/definitions/common.properties/temperature"
                            }
                        }
                    }
                }
            },
            "payload.ThermostatController.general": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "THERMOSTAT_IS_OFF",
                            "UNSUPPORTED_THERMOSTAT_MODE",
                            "DUAL_SETPOINTS_UNSUPPORTED",
                            "TRIPLE_SETPOINTS_UNSUPPORTED",
                            "UNWILLING_TO_SET_VALUE"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    }
                }
            },
            "payload.ThermostatController.REQUESTED_SETPOINTS_TOO_CLOSE": {
                "type": "object",
                "required": [
                    "type"
                ],
                "additionalProperties": false,
                "properties": {
                    "type": {
                        "enum": [
                            "REQUESTED_SETPOINTS_TOO_CLOSE"
                        ]
                    },
                    "message": {
                        "$ref": "#/definitions/common.properties/message"
                    },
                    "minimumTemperatureDelta": {
                        "$ref": "#/definitions/common.properties/temperature"
                    }
                }
            }
        },
        "ResponseOrStateReport.properties": {
            "name": {
                "enum": [
                    "Response",
                    "StateReport",
                    "ActivationStarted",
                    "DeactivationStarted",
                    "AcceptGrant.Response"
                ]
            },
            "with.payload": {
                "type": "object",
                "required": [
                    "event"
                ],
                "additionalProperties": false,
                "properties": {
                    "context": {
                        "$ref": "#/definitions/common.properties/context"
                    },
                    "event": {
                        "type": "object",
                        "required": [
                            "header",
                            "payload"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "header": {
                                "type": "object",
                                "required": [
                                    "namespace",
                                    "name",
                                    "payloadVersion",
                                    "messageId"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "namespace": {
                                        "enum": [
                                            "Alexa.CameraStreamController",
                                            "Alexa.SceneController"
                                        ]
                                    },
                                    "name": {
                                        "$ref": "#/definitions/ResponseOrStateReport.properties/name"
                                    },
                                    "payloadVersion": {
                                        "$ref": "#/definitions/common.properties/payloadVersion"
                                    },
                                    "messageId": {
                                        "$ref": "#/definitions/common.properties/messageId"
                                    },
                                    "correlationToken": {
                                        "$ref": "#/definitions/common.properties/correlationToken"
                                    }
                                }
                            },
                            "endpoint": {
                                "$ref": "#/definitions/common.properties/endpoint"
                            },
                            "payload": {
                                "oneOf": [
                                    {
                                        "$ref": "#/definitions/common.properties/payload/cameraStreams"
                                    },
                                    {
                                        "$ref": "#/definitions/common.properties/payload/sceneActivationDeactivation"
                                    }
                                ],
                                "minProperties": 1
                            }
                        }
                    }
                }
            },
            "without.payload": {
                "type": "object",
                "required": [
                    "event"
                ],
                "additionalProperties": false,
                "properties": {
                    "context": {
                        "$ref": "#/definitions/common.properties/context"
                    },
                    "event": {
                        "type": "object",
                        "required": [
                            "header",
                            "payload"
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "header": {
                                "type": "object",
                                "required": [
                                    "namespace",
                                    "name",
                                    "payloadVersion",
                                    "messageId"
                                ],
                                "additionalProperties": false,
                                "properties": {
                                    "namespace": {
                                        "enum": [
                                            "Alexa",
                                            "Alexa.Authorization"
                                        ]
                                    },
                                    "name": {
                                        "$ref": "#/definitions/ResponseOrStateReport.properties/name"
                                    },
                                    "payloadVersion": {
                                        "$ref": "#/definitions/common.properties/payloadVersion"
                                    },
                                    "messageId": {
                                        "$ref": "#/definitions/common.properties/messageId"
                                    },
                                    "correlationToken": {
                                        "$ref": "#/definitions/common.properties/correlationToken"
                                    }
                                }
                            },
                            "endpoint": {
                                "$ref": "#/definitions/common.properties/endpoint"
                            },
                            "payload": {
                                "type": "object",
                                "additionalProperties": false,
                                "properties": {},
                                "maxProperties": 0
                            }
                        }
                    }
                }
            }
        },
        "ChangeReport.properties": {
            "name": {
                "enum": [
                    "ChangeReport"
                ]
            }
        },
        "DeferredResponse.properties": {
            "name": {
                "enum": [
                    "DeferredResponse"
                ]
            }
        },
        "Discover.Response.properties": {
            "name": {
                "enum": [
                    "Discover.Response"
                ]
            }
        }
    },
    "oneOf": [
        {
            "description": "An ErrorResponse message",
            "type": "object",
            "required": [
                "event"
            ],
            "additionalProperties": false,
            "properties": {
                "context": {
                    "$ref": "#/definitions/common.properties/context"
                },
                "event": {
                    "type": "object",
                    "required": [
                        "header",
                        "payload"
                    ],
                    "oneOf": [
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.general"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.general"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.general"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.general.VALUE_OUT_OF_RANGE"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.general"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.general.NOT_SUPPORTED_IN_CURRENT_MODE"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.general"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.general.ENDPOINT_LOW_POWER"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.ThermostatController"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.ThermostatController.general"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.general"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.general.TEMPERATURE_VALUE_OUT_OF_RANGE"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.ThermostatController"
                                },
                                "endpoint": {
                                    "$ref": "#/definitions/common.properties/endpoint"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.ThermostatController.REQUESTED_SETPOINTS_TOO_CLOSE"
                                }
                            }
                        },
                        {
                            "additionalProperties": false,
                            "properties": {
                                "header": {
                                    "$ref": "#/definitions/ErrorResponse.properties/header.Authorization"
                                },
                                "payload": {
                                    "$ref": "#/definitions/ErrorResponse.properties/payload.general"
                                }
                            }
                        }
                    ]
                }
            }
        },
        {
            "description": "A Response or StateReport message",
            "oneOf": [
                {
                    "$ref": "#/definitions/ResponseOrStateReport.properties/with.payload"
                },
                {
                    "$ref": "#/definitions/ResponseOrStateReport.properties/without.payload"
                }
            ]
        },
        {
            "description": "A ChangeReport message",
            "type": "object",
            "required": [
                "event"
            ],
            "additionalProperties": false,
            "properties": {
                "context": {
                    "$ref": "#/definitions/common.properties/context"
                },
                "event": {
                    "type": "object",
                    "required": [
                        "header",
                        "payload"
                    ],
                    "additionalProperties": false,
                    "properties": {
                        "header": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "payloadVersion",
                                "messageId"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa"
                                    ]
                                },
                                "name": {
                                    "$ref": "#/definitions/ChangeReport.properties/name"
                                },
                                "payloadVersion": {
                                    "$ref": "#/definitions/common.properties/payloadVersion"
                                },
                                "messageId": {
                                    "$ref": "#/definitions/common.properties/messageId"
                                },
                                "correlationToken": {
                                    "$ref": "#/definitions/common.properties/correlationToken"
                                }
                            }
                        },
                        "endpoint": {
                            "$ref": "#/definitions/common.properties/endpoint"
                        },
                        "payload": {
                            "type": "object",
                            "required": [
                                "change"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "change": {
                                    "type": "object",
                                    "required": [
                                        "cause",
                                        "properties"
                                    ],
                                    "additionalProperties": false,
                                    "properties": {
                                        "cause": {
                                            "$ref": "#/definitions/common.properties/cause"
                                        },
                                        "properties": {
                                            "$ref": "#/definitions/common.properties/state.properties"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        {
            "description": "A DeferredResponse message",
            "type": "object",
            "required": [
                "event"
            ],
            "additionalProperties": false,
            "properties": {
                "event": {
                    "type": "object",
                    "required": [
                        "header",
                        "payload"
                    ],
                    "additionalProperties": false,
                    "properties": {
                        "header": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "payloadVersion",
                                "messageId"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa"
                                    ]
                                },
                                "name": {
                                    "$ref": "#/definitions/DeferredResponse.properties/name"
                                },
                                "payloadVersion": {
                                    "$ref": "#/definitions/common.properties/payloadVersion"
                                },
                                "messageId": {
                                    "$ref": "#/definitions/common.properties/messageId"
                                },
                                "correlationToken": {
                                    "$ref": "#/definitions/common.properties/correlationToken"
                                }
                            }
                        },
                        "payload": {
                            "type": "object",
                            "additionalProperties": false,
                            "properties": {
                                "estimatedDeferralInSeconds": {
                                    "type": "integer",
                                    "minimum": 0
                                }
                            }
                        }
                    }
                }
            }
        },
        {
            "description": "A Discover.Response message",
            "type": "object",
            "required": [
                "event"
            ],
            "additionalProperties": false,
            "properties": {
                "event": {
                    "type": "object",
                    "required": [
                        "header",
                        "payload"
                    ],
                    "additionalProperties": false,
                    "properties": {
                        "header": {
                            "type": "object",
                            "required": [
                                "namespace",
                                "name",
                                "payloadVersion",
                                "messageId"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "namespace": {
                                    "enum": [
                                        "Alexa.Discovery"
                                    ]
                                },
                                "name": {
                                    "$ref": "#/definitions/Discover.Response.properties/name"
                                },
                                "payloadVersion": {
                                    "$ref": "#/definitions/common.properties/payloadVersion"
                                },
                                "messageId": {
                                    "$ref": "#/definitions/common.properties/messageId"
                                }
                            }
                        },
                        "payload": {
                            "type": "object",
                            "required": [
                                "endpoints"
                            ],
                            "additionalProperties": false,
                            "properties": {
                                "endpoints": {
                                    "type": "array",
                                    "uniqueItems": true,
                                    "maxItems": 300,
                                    "items": {
                                        "type": "object",
                                        "required": [
                                            "endpointId",
                                            "manufacturerName",
                                            "friendlyName",
                                            "description",
                                            "displayCategories",
                                            "capabilities"
                                        ],
                                        "properties": {
                                            "endpointId": {
                                                "$ref": "#/definitions/common.properties/endpointId"
                                            },
                                            "manufacturerName": {
                                                "type": "string",
                                                "minLength": 1,
                                                "maxLength": 128
                                            },
                                            "friendlyName": {
                                                "type": "string",
                                                "minLength": 1,
                                                "maxLength": 128
                                            },
                                            "description": {
                                                "type": "string",
                                                "minLength": 1,
                                                "maxLength": 128
                                            },
                                            "displayCategories": {
                                                "type": "array",
                                                "minItems": 1,
                                                "uniqueItems": true,
                                                "items": {
                                                    "enum": [
                                                        "LIGHT",
                                                        "SMARTPLUG",
                                                        "SWITCH",
                                                        "CAMERA",
                                                        "DOOR",
                                                        "THERMOSTAT",
                                                        "TEMPERATURE_SENSOR",
                                                        "SMARTLOCK",
                                                        "SCENE_TRIGGER",
                                                        "ACTIVITY_TRIGGER",
                                                        "OTHER"
                                                    ]
                                                }
                                            },
                                            "cookie": {
                                                "type": "object",
                                                "additionalProperties": {
                                                    "type": "string"
                                                }
                                            },
                                            "capabilities": {
                                                "type": "array",
                                                "minItems": 1,
                                                "uniqueItems": true,
                                                "additionalItems": false,
                                                "items": {
                                                    "anyOf": [
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/Alexa/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/EndpointHealth/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/PowerController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/PowerLevelController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/PercentageController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/BrightnessController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/CameraStreamController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/ColorController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/ColorTemperatureController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/LockController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/ThermostatController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/TemperatureSensor/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/ChannelController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/InputController/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/Speaker/capabilities"
                                                        },
                                                        {
                                                            "$ref": "#/definitions/common.properties/interfaces/SceneController/capabilities"
                                                        }
                                                    ]
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    ]
}
`
