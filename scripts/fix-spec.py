#!/usr/bin/env python3

import yaml

SPEC_PATH = 'api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():
        if 'properties' in schema:
            # Remove "null" item from nullable enums
            for prop_name, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]

            for ntype in nullable_types:
                if ntype in schema['properties']:
                    schema['properties'][ntype]['nullable'] = True

            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]

            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    if schema['properties'][ntype]['format'] == 'binary':
                        schema['properties'][ntype].pop('nullable')

            # Fix non-required fields
            non_required_types = [
                'virtualmachine_count',
                'devicetype_count',
                'device_count',
                'site_count',
                'provider_count',
                'prefix_count',
                'cluster_count',
                'vlan_count',
                'vrf_count',
                'nat_outside',
            ]

            # if name ends with "Request"
            if name.endswith('Request'):
                non_required_types.append('name')
                non_required_types.append('slug')

            if name == 'BriefVLANRequest':
                non_required_types.append('vid')

            if name == 'BriefIPAddressRequest':
                non_required_types.append('address')

            # remove non-required fields
            if 'required' in schema:
                # foreach all properties
                for name, prop in schema['properties'].items():
                    # if schema says it's required
                    if name in schema['required']:
                        # if it's nullable or in non_required_types
                        if prop.get('nullable') == True or name in non_required_types:
                            # remove from required, because netbox may not send it
                            schema['required'].remove(name)

# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
