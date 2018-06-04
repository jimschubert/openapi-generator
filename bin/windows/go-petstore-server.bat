set executable=.\modules\openapi-generator-cli\build\current\openapi-generator-cli.jar

If Not Exist %executable% (
  mvn clean package
)

REM set JAVA_OPTS=%JAVA_OPTS% -Xmx1024M -DloggerPath=conf/log4j.properties
set ags=generate -i modules\openapi-generator\src\test\resources\2_0\petstore.yaml -g go-server -o samples\server\petstore\go-api-server -DpackageName=petstoreserver  -Dservice

java %JAVA_OPTS% -jar %executable% %ags%
