#!/bin/bash

padsp signalgen -t 100m sin 444

input=$(ls *.md | randline)
output=$(ls *.md | randline)

echo "Mime-Version: 1.0
Content-Transfer-Encoding: quoted-printable
Sysmsg: You are an expert technical writer and software architect. 
    Merge the input documents into the output document.  Maintain the
    focus of the output document.  Reconcile any conflicts between the
    input documents.  The output document should be a coherent whole.
In: 
    v2/doc/$input
    v2/doc/$output
Out: v2/doc/$output

Merge $input into $output without changing the focus of $output.
Reconcile any conflicts between the two documents.  The output document
should be a coherent whole." > ../../.aidda/prompt

grok aidda commit prompt

padsp signalgen -t 100m sin 500
