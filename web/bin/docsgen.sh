#!/bin/sh

if which godocdown >/dev/null
then

  echo "---"                  >  content/doc/index.md
  echo "title: Documentation" >> content/doc/index.md
  echo "---"                  >> content/doc/index.md
  echo "## Packages index"  >> content/doc/index.md

  for pkg in cmd pkg/config pkg/file pkg/metadata pkg/operation pkg/resource pkg/term
  do
    echo "---"         >  content/doc/packages/$pkg.md
    echo "title: $pkg" >> content/doc/packages/$pkg.md
    echo "---"         >> content/doc/packages/$pkg.md

    echo "\`\`\`go"                               >> content/doc/packages/$pkg.md
    echo "import \"github.com/mod/kaigara/$pkg\"" >> content/doc/packages/$pkg.md
    echo "\`\`\`"                                 >> content/doc/packages/$pkg.md

    godocdown github.com/mod/kaigara/$pkg | sed "/^# /d" | sed "/^--$/d" | sed "/^\s\s\s\simport/d" >> content/doc/packages/$pkg.md
    echo "* [$pkg](/doc/packages/$pkg)" >> content/doc/index.md
  done

else
  echo "Please install godocdown!"
fi
