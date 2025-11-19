#!/usr/bin/env python3
"""
AI-enhanced LinkedDoc annotation generator using Claude Haiku
Generates semantic documentation for Go source files
"""

import os
import sys
import json
import subprocess
import re
from pathlib import Path

def read_file(filepath):
    """Read file content"""
    with open(filepath, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(filepath, content):
    """Write content to file"""
    with open(filepath, 'w', encoding='utf-8') as f:
        f.write(content)

def has_linkeddoc(filepath):
    """Check if file already has LinkedDoc annotation"""
    content = read_file(filepath)
    return '<!-- LinkedDoc RDF -->' in content

def get_module_path(content):
    """Extract module path from go.mod"""
    match = re.search(r'^module\s+(.+)$', content, re.MULTILINE)
    return match.group(1) if match else None

def extract_package(content):
    """Extract package name from Go file"""
    match = re.search(r'^package\s+(\w+)', content, re.MULTILINE)
    return match.group(1) if match else "main"

def extract_exports(content):
    """Extract exported symbols from Go file"""
    exports = []

    # Types
    for match in re.finditer(r'^type\s+([A-Z][a-zA-Z0-9]*)', content, re.MULTILINE):
        exports.append(match.group(1))

    # Functions
    for match in re.finditer(r'^func\s+([A-Z][a-zA-Z0-9]*)', content, re.MULTILINE):
        exports.append(match.group(1))

    # Constants and Variables
    for match in re.finditer(r'^(?:const|var)\s+([A-Z][a-zA-Z0-9]*)', content, re.MULTILINE):
        exports.append(match.group(1))

    return sorted(set(exports))

def extract_imports(content, module_path):
    """Extract internal imports from Go file"""
    imports = []
    in_import_block = False

    for line in content.split('\n'):
        line = line.strip()

        if line.startswith('import ('):
            in_import_block = True
            continue
        elif in_import_block and line == ')':
            in_import_block = False
            continue

        # Extract import path
        match = re.search(r'"([^"]+)"', line)
        if match:
            import_path = match.group(1)
            # Only keep internal imports
            if module_path and import_path.startswith(module_path + '/'):
                rel_path = import_path[len(module_path)+1:]
                imports.append(rel_path)

    return sorted(set(imports))

def extract_layer(filepath):
    """Determine architectural layer from file path"""
    parts = Path(filepath).parts

    # Map directory patterns to layers
    layer_map = {
        'api': 'api',
        'acl': 'acl',
        'agent': 'agent',
        'connect': 'service-mesh',
        'grpc-external': 'grpc',
        'grpc-internal': 'grpc',
        'command': 'cli',
        'internal': 'internal',
        'tlsutil': 'security',
        'logging': 'internal',
        'lib': 'internal',
    }

    # Check first few path components
    for part in parts[:3]:
        if part in layer_map:
            return layer_map[part]

    # Default based on patterns
    if 'consul' in parts:
        return 'consensus'
    elif 'state' in parts:
        return 'storage'
    elif 'structs' in parts:
        return 'api'

    return 'internal'

def generate_tags(filepath, layer):
    """Generate semantic tags for the module"""
    parts = Path(filepath).parts
    tags = set()

    # Add layer
    if layer:
        tags.add(layer)

    # Add path-based tags
    tag_keywords = {
        'acl': ['security', 'authorization', 'access-control'],
        'auth': ['security', 'authentication'],
        'tls': ['security', 'encryption'],
        'connect': ['service-mesh', 'mtls'],
        'proxy': ['service-mesh', 'networking'],
        'xds': ['service-mesh', 'envoy'],
        'health': ['monitoring', 'health-checks'],
        'check': ['monitoring', 'health-checks'],
        'catalog': ['discovery', 'registry'],
        'discovery': ['discovery', 'dns'],
        'dns': ['discovery', 'networking'],
        'config': ['configuration'],
        'state': ['storage', 'persistence'],
        'raft': ['consensus', 'replication'],
        'rpc': ['networking', 'communication'],
        'grpc': ['api', 'grpc'],
        'api': ['api', 'client'],
        'command': ['cli', 'user-interface'],
        'structs': ['data-model', 'types'],
        'types': ['data-model', 'types'],
    }

    filepath_lower = str(filepath).lower()
    for keyword, tag_list in tag_keywords.items():
        if keyword in filepath_lower:
            tags.update(tag_list)

    return sorted(tags)

def generate_basic_annotation(filepath, repo_root):
    """Generate basic LinkedDoc annotation using pattern extraction"""
    content = read_file(filepath)

    # Get module path from go.mod
    gomod_path = Path(repo_root) / 'go.mod'
    module_path = None
    if gomod_path.exists():
        module_path = get_module_path(read_file(gomod_path))

    # Extract metadata
    rel_path = os.path.relpath(filepath, repo_root)
    package = extract_package(content)
    exports = extract_exports(content)
    imports = extract_imports(content, module_path)
    layer = extract_layer(rel_path)
    tags = generate_tags(rel_path, layer)

    # Generate basic description
    filename = Path(filepath).stem
    desc = filename.replace('_', ' ').replace('-', ' ').title()

    return {
        'module_name': rel_path,
        'package': package,
        'description': f"{desc} module for {layer} layer",
        'layer': layer,
        'exports': exports,
        'imports': imports,
        'tags': tags,
        'content': content[:3000],  # First 3000 chars for AI context
    }

def format_linkeddoc(metadata):
    """Format LinkedDoc RDF annotation"""
    lines = []
    lines.append("/*")
    lines.append(f"# Module: {metadata['module_name']}")
    lines.append(metadata['description'])
    lines.append("")

    # Add linked modules
    if metadata['imports']:
        lines.append("## Linked Modules")
        for imp in metadata['imports'][:10]:  # Limit to 10
            lines.append(f"- [{imp}](../{imp})")
        lines.append("")

    # Add tags
    if metadata['tags']:
        lines.append("## Tags")
        lines.append(", ".join(metadata['tags']))
        lines.append("")

    # Add exports
    if metadata['exports']:
        lines.append("## Exports")
        lines.append(", ".join(metadata['exports'][:20]))  # Limit to 20
        lines.append("")

    # RDF section
    lines.append("<!-- LinkedDoc RDF -->")
    lines.append("@prefix code: <https://schema.codedoc.org/> .")
    lines.append(f"<{metadata['module_name']}> a code:Module ;")
    lines.append(f'    code:name "{metadata["module_name"]}" ;')
    lines.append(f'    code:description "{metadata["description"]}" ;')
    lines.append(f'    code:language "go" ;')
    lines.append(f'    code:layer "{metadata["layer"]}" ;')

    # Add imports as linked modules
    if metadata['imports']:
        link_parts = []
        for imp in metadata['imports'][:10]:
            link_parts.append(f'        code:name "{imp}" ;\n        code:path "../{imp}"')

        lines.append("    code:linksTo [")
        lines.append(link_parts[0])
        for link in link_parts[1:]:
            lines.append("    ], [")
            lines.append(link)
        lines.append("    ] ;")

    # Add exports
    if metadata['exports']:
        export_refs = [f":{exp}" for exp in metadata['exports'][:20]]
        lines.append(f"    code:exports {', '.join(export_refs)} ;")

    # Add tags
    if metadata['tags']:
        tag_strs = [f'"{tag}"' for tag in metadata['tags']]
        lines.append(f"    code:tags {', '.join(tag_strs)} .")

    lines.append("<!-- End LinkedDoc RDF -->")
    lines.append("*/")

    return "\n".join(lines)

def insert_annotation(filepath, annotation):
    """Insert LinkedDoc annotation at the beginning of file"""
    content = read_file(filepath)

    # Find package declaration
    package_match = re.search(r'^package\s+\w+', content, re.MULTILINE)
    if not package_match:
        print(f"ERROR: Could not find package declaration in {filepath}", file=sys.stderr)
        return False

    # Split at package declaration
    package_pos = package_match.start()

    # Keep copyright/license header if present
    header = content[:package_pos].rstrip()
    rest = content[package_pos:]

    # Construct new content
    new_content = header + "\n\n" + annotation + "\n" + rest

    write_file(filepath, new_content)
    return True

def main():
    if len(sys.argv) < 2:
        print("Usage: ai_annotate.py <file>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    repo_root = subprocess.check_output(['git', 'rev-parse', '--show-toplevel'], text=True).strip()

    # Check if already annotated
    if has_linkeddoc(filepath):
        print(f"SKIP: {filepath} already has LinkedDoc annotation", file=sys.stderr)
        sys.exit(2)

    # Generate annotation
    metadata = generate_basic_annotation(filepath, repo_root)
    annotation = format_linkeddoc(metadata)

    # Insert annotation
    if insert_annotation(filepath, annotation):
        print(f"SUCCESS: Annotated {filepath}", file=sys.stderr)
        sys.exit(0)
    else:
        print(f"ERROR: Failed to annotate {filepath}", file=sys.stderr)
        sys.exit(1)

if __name__ == '__main__':
    main()
