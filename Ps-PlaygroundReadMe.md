Welcome to the Polyscripting Playground


From within the Docker container here are the steps to running polyscripted php.

Install standard php configured to support Polyscripting:

# ./build-php.sh

Next, to build a scrambled version of php:

# ./build-scrambled.sh

Or if you want to generate expected outputs for the .php files in tests:

# ./build-scrambled.sh -t

Note: You can build a newly scrambled version of php at any time with
the build-scrambled script, but option -t will only generate expected test
outputs accurately if you are building from the standard php.

Now to the fun stuff.
Notice that if you try to run a php program, you will be hit with an error.

Parse error: syntax error, unexpected 'as'
Because php no longer recognizes its standard keywords.

To transform your php file to polyscripted use ./transformer :

# ./transformer -f [filename]

Here's an example:

# ./transformer -f tests/smallWorld.php

This will generate a new file with an appended ps- to the file name.
If you'd like to overwrite the original file use the option -replace

# ./transformer -f tests/smallWorld.php -replace=true

To run the newly created file use /polyscripted-php/bin/php where you would
typically use the php command.

# /polyscripted-php/bin/php tests/ps-smallWorld.php

This will use polyscripted php to interpret your  polyscripted file.

If you ran the -t option when scrambling your build you can see the difference
in the outputs of using standard php and scrambled php by using the command:

# diff <(/polscripted-php/bin/php tests/ps-smallWorld.php) expected_smallWorld.php


Feel free to try it out with your own php files. There's also a small program with an eval vulnerability in tests/evalExploit if you're unfamiliar with code injection attacks, then run this with standard php first. Then scramble it up, and see what you can't do.

Note: You can build a newly scrambled version of php at any time with
the build-scrambled script, but option -t will only generate expected test
outputs accurately if you are building from the standard php.

To revert back to standard php use the command:

# ./resetPhp/reset-php.sh -revert
