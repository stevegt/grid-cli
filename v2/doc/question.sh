#!/bin/bash

set -x

# padsp signalgen -t 100m sin 444

pfile=/tmp/$$.prompt
qfile=/tmp/qfile.txt
touch $qfile

while true
do
    chatfile=/tmp/$$.chat
    rm -f $chatfile

    grok aidda commit 
    padsp signalgen -t 100m sin 600

    inputs=$(ls -rt *.md)
    inputs=$(echo $inputs | sed 's/ /,/g')

    sysmsg="You are an expert technical writer and software architect. 
        We are writing a book about PromiseGrid.  Each of the input files
        is a chapter in the book."
    prompt="Examine the chapters.  Ask me one question about one of the
        chapters.  Pay particular attention to resolving ambiguities
        and inconsistencies. The first line of your question should be
        the heading 'File:' followed only by the filename of the
        chapter you are asking about.  The rest of the question should
        be the question itself."
    echo $prompt | grok chat $chatfile -s "$sysmsg" -i $inputs > $qfile

    echo "Waiting for answer from $qfile..."
    padsp signalgen -t 100m sin 700
    inotifywait -e modify $qfile
    padsp signalgen -t 100m sin 444

    fn=$(head -n 1 $qfile | sed 's/File: //')
    if [ -z "$fn" ]; then
        echo "No filename found."
        padsp signalgen -t 100m sin 300
        continue
    fi

    echo "Modify $fn to answer the following question:" > $pfile
    echo >> $pfile
    tail -n +1 $qfile >> $pfile

    grok chat $chatfile -s "$sysmsg" -i $fn -o $fn < $pfile
    padsp signalgen -t 100m sin 500
    inotifywait -e modify $fn
done



