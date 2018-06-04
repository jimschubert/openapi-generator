set executable=.\modules\openapi-generator-cli\build\current\openapi-generator-cli.jar

If Not Exist %executable% (
  mvn clean package
)

REM set JAVA_OPTS=%JAVA_OPTS% -Xmx1024M
set ags=generate -i modules\openapi-generator\src\test\resources\2_0\petstore.yaml -g java-play-framework -o samples\server\petstore\java-play-framework -DhideGenerationTimestamp=true

java %JAVA_OPTS% -jar %executable% %ags%
