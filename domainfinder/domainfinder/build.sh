#!/bin/bash

echo Start building

# domainfinder
echo Building domainfinder...
go build -o domainfinder

# synonyms
echo Building synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms

# available
echo Building available...
cd ../available
go build -o ../domainfinder/lib/available

# sprinkle
echo Building sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle

# coolify
echo Building coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolify

# domainify
echo Building domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainify

cd ../domainfinder
echo Done.