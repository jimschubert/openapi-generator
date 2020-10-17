/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using JsonSubTypes;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// Pig
    /// </summary>
    [DataContract(Name = "Pig")]
    public partial class Pig : AbstractOpenAPISchema, IEquatable<Pig>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="Pig" /> class.
        /// </summary>
        public Pig()
        {
            this.IsNullable = true;
            this.SchemaType= "oneOf";
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="Pig" /> class
        /// with the <see cref="BasquePig" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of BasquePig.</param>
        public Pig(BasquePig actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="Pig" /> class
        /// with the <see cref="DanishPig" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of DanishPig.</param>
        public Pig(DanishPig actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }


        private Object _actualInstance;

        /// <summary>
        /// Gets or Sets ActualInstance
        /// </summary>
        public override Object ActualInstance
        {
            get
            {
                return _actualInstance;
            }
            set
            {
                if (value.GetType() == typeof(BasquePig))
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(DanishPig))
                {
                    this._actualInstance = value;
                }
                else
                {
                    throw new ArgumentException("Invalid instance found. Must be the following types: BasquePig, DanishPig");
                }
            }
        }

        /// <summary>
        /// Get the actual instance of `BasquePig`. If the actual instanct is not `BasquePig`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of BasquePig</returns>
        public BasquePig GetBasquePig()
        {
            return (BasquePig)this.ActualInstance;
        }

        /// <summary>
        /// Get the actual instance of `DanishPig`. If the actual instanct is not `DanishPig`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of DanishPig</returns>
        public DanishPig GetDanishPig()
        {
            return (DanishPig)this.ActualInstance;
        }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class Pig {\n");
            sb.Append("  ActualInstance: ").Append(this.ActualInstance).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public override string ToJson()
        {
            return JsonConvert.SerializeObject(this.ActualInstance, _serializerSettings);
        }

        /// <summary>
        /// Converts the JSON string into an instance of Pig
        /// </summary>
        /// <param name="jsonString">JSON string</param>
        /// <returns>An instance of Pig</returns>
        public static Pig FromJson(string jsonString)
        {
            Pig newPig = new Pig();

            string discriminatorValue = JObject.Parse(jsonString)["className"].ToString();
            switch (discriminatorValue)
            {
                case "BasquePig":
                    newPig.ActualInstance = JsonConvert.DeserializeObject<BasquePig>(jsonString, newPig._serializerSettings);
                    return newPig;
                case "DanishPig":
                    newPig.ActualInstance = JsonConvert.DeserializeObject<DanishPig>(jsonString, newPig._serializerSettings);
                    return newPig;
                default:
                    System.Diagnostics.Debug.WriteLine(String.Format("Failed to lookup discriminator value `%s` for Pig. Possible values: BasquePig DanishPig", discriminatorValue));
                    break;
            }

            int match = 0;
            List<string> matchedTypes = new List<string>();

            try
            {
                newPig.ActualInstance = JsonConvert.DeserializeObject<BasquePig>(jsonString, newPig._serializerSettings);
                matchedTypes.Add("BasquePig");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(String.Format("Failed to deserialize `%s` into BasquePig: %s", jsonString, exception.ToString()));
            }

            try
            {
                newPig.ActualInstance = JsonConvert.DeserializeObject<DanishPig>(jsonString, newPig._serializerSettings);
                matchedTypes.Add("DanishPig");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(String.Format("Failed to deserialize `%s` into DanishPig: %s", jsonString, exception.ToString()));
            }

            if (match == 0)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` cannot be deserialized into any schema defined.");
            }
            else if (match > 1)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` incorrectly matches more than one schema (should be exactly one match): " + matchedTypes);
            }
            
            // deserialization is considered successful at this point if no exception has been thrown.
            return newPig;
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input as Pig).AreEqual;
        }

        /// <summary>
        /// Returns true if Pig instances are equal
        /// </summary>
        /// <param name="input">Instance of Pig to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(Pig input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input).AreEqual;
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.ActualInstance != null)
                    hashCode = hashCode * 59 + this.ActualInstance.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
