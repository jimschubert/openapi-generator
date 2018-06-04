set executable=.\modules\openapi-generator-cli\build\current\openapi-generator-cli.jar

If Not Exist %executable% (
  mvn clean package
)

REM set JAVA_OPTS=%JAVA_OPTS% -Xmx1024M -DloggerPath=conf/log4j.properties
set ags=generate  --artifact-id "scala-akka-petstore-client" -i modules\openapi-generator\src\test\resources\2_0\petstore.yaml -g scala-akka -o samples\client\petstore\akka-scala

java %JAVA_OPTS% -jar %executable% %ags%
