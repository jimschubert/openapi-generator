/*
 * Copyright 2018 OpenAPI-Generator Contributors (https://openapi-generator.tech)
 * Copyright 2018 SmartBear Software
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.openapitools.codegen.cmd;

import io.airlift.airline.Command;
import io.airlift.airline.Option;

import io.swagger.parser.OpenAPIParser;
import io.swagger.v3.oas.models.Components;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.Paths;
import io.swagger.v3.oas.models.media.ComposedSchema;
import io.swagger.v3.oas.models.media.Schema;
import io.swagger.v3.oas.models.security.SecurityScheme;
import io.swagger.v3.parser.core.models.SwaggerParseResult;
import org.apache.commons.lang3.text.WordUtils;
import org.openapitools.codegen.utils.ModelUtils;
import org.openapitools.codegen.validation.ValidationResult;
import org.openapitools.codegen.validations.OpenApiEvaluator;
import org.openapitools.codegen.validations.OpenApiParameterValidations;
import org.openapitools.codegen.validations.OpenApiSecuritySchemeValidations;
import org.openapitools.codegen.validations.RuleConfiguration;

import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

@Command(name = "validate", description = "Validate specification")
public class Validate implements Runnable {

    @Option(name = {"-i", "--input-spec"}, title = "spec file", required = true,
            description = "location of the OpenAPI spec, as URL or file (required)")
    private String spec;

    @Option(name = { "--recommend"}, title = "recommend spec improvements")
    private Boolean recommend;

    @Override
    public void run() {
        System.out.println("Validating spec (" + spec + ")");

        SwaggerParseResult result = new OpenAPIParser().readLocation(spec, null, null);
        List<String> messageList = result.getMessages();
        Set<String> errors = new HashSet<>(messageList);
        Set<String> warnings = new HashSet<>();

        StringBuilder sb = new StringBuilder();
        OpenAPI specification = result.getOpenAPI();

        RuleConfiguration ruleConfiguration = new RuleConfiguration();
        ruleConfiguration.setEnableRecommendations(recommend != null ? recommend : true);

        OpenApiEvaluator evaluator = new OpenApiEvaluator(ruleConfiguration);
        ValidationResult validationResult = evaluator.validate(specification);

        // TODO: We could also provide description here along with getMessage. getMessage is either a "generic" message or specific (e.g. Model 'Cat' has issues).
        //       This would require that we parse the messageList coming from swagger-parser into a better structure.
        validationResult.getWarnings().forEach(invalid -> warnings.add(invalid.getMessage()));
        validationResult.getErrors().forEach(invalid -> errors.add(invalid.getMessage()));

        if (!errors.isEmpty()) {
            sb.append("Errors:").append(System.lineSeparator());
            errors.forEach(msg ->
                    sb.append("\t- ").append(WordUtils.wrap(msg, 90).replace(System.lineSeparator(), System.lineSeparator() + "\t  ")).append(System.lineSeparator())
            );
        }

        if (!warnings.isEmpty()) {
            sb.append("Warnings: ").append(System.lineSeparator());
            warnings.forEach(msg ->
                    sb.append("\t- ").append(WordUtils.wrap(msg, 90).replace(System.lineSeparator(), System.lineSeparator() + "\t  ")).append(System.lineSeparator())
            );
        }

        if (!errors.isEmpty()) {
            sb.append(System.lineSeparator());
            sb.append("[error] Spec has ").append(errors.size()).append(" errors.");
            System.err.println(sb.toString());
            System.exit(1);
        } else if (!warnings.isEmpty()) {
            sb.append(System.lineSeparator());
            sb.append("[info] Spec has ").append(warnings.size()).append(" recommendation(s).");
        } else {
            // we say "issues" here rather than "errors" to account for both errors and issues.
            sb.append("No validation issues detected.");
        }

        System.out.println(sb.toString());
    }
}
