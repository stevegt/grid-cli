#!/bin/bash

# padsp signalgen -t 100m sin 444

output=$1

if [ -z "$output" ]; then
    output=$(ls *.md | randline)
fi

echo "Mime-Version: 1.0
Content-Transfer-Encoding: quoted-printable
Sysmsg: You are an expert technical writer and software architect. 
    You must make changes to the output document.  Do not leave the
    output document unchanged.  Reconcile the output document with all
    input documents, removing or correcting any conflicts.  The output
    document should be a coherent whole.  Do not change the central
    focus of the output document.
In: v2/doc/
Out: v2/doc/$output

Revise $output; you must make changes.  Do not change the focus of
$output.  Reconcile or remove any information in $output that
conflicts with the input documents.  The output document should be a
coherent whole.

If the input or output documents contain any open questions, answer
them in the output document.  

Add new open questions to the output document.
" > ../../.aidda/prompt

grok aidda commit prompt

# padsp signalgen -t 100m sin 500
