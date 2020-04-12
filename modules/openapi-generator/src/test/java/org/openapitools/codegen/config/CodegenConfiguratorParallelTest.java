package org.openapitools.codegen.config;

import java.util.stream.IntStream;

import org.openapitools.codegen.config.CodegenConfigurator;
import org.testng.annotations.Test;

public class CodegenConfiguratorParallelTest {

    @Test
    public void test() {
        final String file = "src/test/resources/sampleConfig.json";

        IntStream.range(0, 100).parallel()
                .forEach(i -> {
                    System.out.println(i + ": " + Thread.currentThread().getId());
                    CodegenConfigurator.fromFile(file);
                });
    }
}