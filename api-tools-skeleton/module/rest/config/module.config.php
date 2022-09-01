<?php
return [
    'router' => [
        'routes' => [
            'rest.rest.users' => [
                'type' => 'Segment',
                'options' => [
                    'route' => '/users[/:users_id]',
                    'defaults' => [
                        'controller' => 'rest\\V1\\Rest\\Users\\Controller',
                    ],
                ],
            ],
        ],
    ],
    'api-tools-versioning' => [
        'uri' => [
            0 => 'rest.rest.users',
        ],
    ],
    'api-tools-rest' => [
        'rest\\V1\\Rest\\Users\\Controller' => [
            'listener' => 'rest\\V1\\Rest\\Users\\UsersResource',
            'route_name' => 'rest.rest.users',
            'route_identifier_name' => 'users_id',
            'collection_name' => 'users',
            'entity_http_methods' => [
                0 => 'GET',
                1 => 'PATCH',
                2 => 'PUT',
                3 => 'DELETE',
                4 => 'POST',
            ],
            'collection_http_methods' => [
                0 => 'GET',
                1 => 'POST',
            ],
            'collection_query_whitelist' => [],
            'page_size' => 25,
            'page_size_param' => null,
            'entity_class' => \rest\V1\Rest\Users\UsersEntity::class,
            'collection_class' => \rest\V1\Rest\Users\UsersCollection::class,
            'service_name' => 'users',
        ],
    ],
    'api-tools-content-negotiation' => [
        'controllers' => [
            'rest\\V1\\Rest\\Users\\Controller' => 'HalJson',
        ],
        'accept_whitelist' => [
            'rest\\V1\\Rest\\Users\\Controller' => [
                0 => 'application/vnd.rest.v1+json',
                1 => 'application/hal+json',
                2 => 'application/json',
            ],
        ],
        'content_type_whitelist' => [
            'rest\\V1\\Rest\\Users\\Controller' => [
                0 => 'application/vnd.rest.v1+json',
                1 => 'application/json',
            ],
        ],
    ],
    'api-tools-hal' => [
        'metadata_map' => [
            \rest\V1\Rest\Users\UsersEntity::class => [
                'entity_identifier_name' => 'id',
                'route_name' => 'rest.rest.users',
                'route_identifier_name' => 'users_id',
                'hydrator' => \Laminas\Hydrator\ArraySerializable::class,
            ],
            \rest\V1\Rest\Users\UsersCollection::class => [
                'entity_identifier_name' => 'id',
                'route_name' => 'rest.rest.users',
                'route_identifier_name' => 'users_id',
                'is_collection' => true,
            ],
        ],
    ],
    'api-tools' => [
        'db-connected' => [
            'rest\\V1\\Rest\\Users\\UsersResource' => [
                'adapter_name' => 'dummy',
                'table_name' => 'users',
                'hydrator_name' => \Laminas\Hydrator\ArraySerializable::class,
                'controller_service_name' => 'rest\\V1\\Rest\\Users\\Controller',
                'entity_identifier_name' => 'id',
                'table_service' => 'rest\\V1\\Rest\\Users\\UsersResource\\Table',
            ],
        ],
    ],
];
