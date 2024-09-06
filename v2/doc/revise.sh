#!/bin/bash

# padsp signalgen -t 100m sin 444

output=$1

if [ -z "$output" ]; then
    output=$(ls *.md | randline)
fi

echo "Mime-Version: 1.0
Content-Transfer-Encoding: quoted-printable
Sysmsg: You are an expert technical writer and software architect. 
    You are writing a book about PromiseGrid.  Each of the input files
    is a chapter in the book.
In: v2/doc/
Out: v2/doc/$output

Examine the chapters.  Look for conflicts, redundancies, and
inconsistencies compared to $output.  Fix $output.
You must make changes to $output.  Do not change the focus of
$output.  The resulting document should be a coherent whole.
" > ../../.aidda/prompt

grok aidda commit prompt

# padsp signalgen -t 100m sin 500
