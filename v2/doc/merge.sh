#!/bin/bash

padsp signalgen -t 100m sin 444

input=$(ls *.md | randline)
output=$(ls *.md | randline)

echo "Mime-Version: 1.0
Content-Transfer-Encoding: quoted-printable
Sysmsg: You are an expert technical writer and software architect. 
    Reconcile the output document with the input document, removing or
    correcting any conflicts between the two documents.  The output
    document should be a coherent whole.  Do not change the central
    focus of the output document.
In: 
    v2/doc/$input
    v2/doc/$output
Out: v2/doc/$output

Reconcile $output with $input.  Do not change the focus of $output.
Reconcile or remove any information in $output that conflicts with $input.
The output document should be a coherent whole.

If the input or output documents contain any open questions, answer
them in the output document.  

Add new open questions to the output document if they are not already
present.
" > ../../.aidda/prompt

I

grok aidda commit prompt

padsp signalgen -t 100m sin 500
