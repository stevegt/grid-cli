#!/bin/bash

set -x

# padsp signalgen -t 100m sin 444

# ensure OLLAMA_HOST is set
if [ -z "$OLLAMA_HOST" ]; then
    echo "set OLLAMA_HOST to the host:port of the OLLAMA server"
    # padsp signalgen -t 100m sin 300
    exit 1
fi

cd ../../
files=$(~/lab/ollama/go-file-similarity-graph/main -closest 2 v2/doc/*.md)
files=$(echo "$files" | grep -v "doc/README.md")
files=$(echo "$files" | tr '\n' ' ')

input1=$(echo "$files" | awk '{print $1}')
input2=$(echo "$files" | awk '{print $2}')
output=$input1

if [ -z "$input2" ]; then
    echo "No similar files found"
    # padsp signalgen -t 100m sin 300
    exit 1
fi

echo "Mime-Version: 1.0
Content-Transfer-Encoding: quoted-printable
Sysmsg: You are an expert technical writer and software architect. 
In: 
    $input1
    $input2
Out: $output

- merge all input files into the output file
- resolve inconsistencies
- remove redundant information
- focus on engineering details
- remove fluff and marketing 
- be concise
" > .aidda/prompt

grok aidda commit prompt

# padsp signalgen -t 100m sin 500
