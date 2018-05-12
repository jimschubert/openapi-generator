/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apiteam@swagger.io
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.OpenAPIPetstore);
  }
}(this, function(expect, OpenAPIPetstore) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new OpenAPIPetstore.OuterComposite();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('OuterComposite', function() {
    it('should create an instance of OuterComposite', function() {
      // uncomment below and update the code to test OuterComposite
      //var instane = new OpenAPIPetstore.OuterComposite();
      //expect(instance).to.be.a(OpenAPIPetstore.OuterComposite);
    });

    it('should have the property myNumber (base name: "my_number")', function() {
      // uncomment below and update the code to test the property myNumber
      //var instane = new OpenAPIPetstore.OuterComposite();
      //expect(instance).to.be();
    });

    it('should have the property myString (base name: "my_string")', function() {
      // uncomment below and update the code to test the property myString
      //var instane = new OpenAPIPetstore.OuterComposite();
      //expect(instance).to.be();
    });

    it('should have the property myBoolean (base name: "my_boolean")', function() {
      // uncomment below and update the code to test the property myBoolean
      //var instane = new OpenAPIPetstore.OuterComposite();
      //expect(instance).to.be();
    });

  });

}));
