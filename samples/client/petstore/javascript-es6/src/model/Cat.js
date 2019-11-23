/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import Animal from './Animal';
import CatAllOf from './CatAllOf';

/**
 * The Cat model module.
 * @module model/Cat
 * @version 1.0.0
 */
class Cat {
    /**
     * Constructs a new <code>Cat</code>.
     * @alias module:model/Cat
     * @extends module:model/Animal
     * @implements module:model/Animal
     * @implements module:model/CatAllOf
     * @param className {String} 
     */
    constructor(className) { 
        Animal.initialize(this, className);CatAllOf.initialize(this);
        Cat.initialize(this, className);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, className) { 
    }

    /**
     * Constructs a <code>Cat</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Cat} obj Optional instance to populate.
     * @return {module:model/Cat} The populated <code>Cat</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Cat();
            Animal.constructFromObject(data, obj);
            Animal.constructFromObject(data, obj);
            CatAllOf.constructFromObject(data, obj);

            if (data.hasOwnProperty('declawed')) {
                obj['declawed'] = ApiClient.convertToType(data['declawed'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} declawed
 */
Cat.prototype['declawed'] = undefined;


// Implement Animal interface:
/**
 * @member {String} className
 */
Animal.prototype['className'] = undefined;
/**
 * @member {String} color
 * @default 'red'
 */
Animal.prototype['color'] = 'red';
// Implement CatAllOf interface:
/**
 * @member {Boolean} declawed
 */
CatAllOf.prototype['declawed'] = undefined;




export default Cat;

