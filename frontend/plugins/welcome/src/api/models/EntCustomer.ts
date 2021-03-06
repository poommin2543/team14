/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCustomerEdges,
    EntCustomerEdgesFromJSON,
    EntCustomerEdgesFromJSONTyped,
    EntCustomerEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCustomer
 */
export interface EntCustomer {
    /**
     * Address holds the value of the "Address" field.
     * @type {string}
     * @memberof EntCustomer
     */
    address?: string;
    /**
     * Customername holds the value of the "Customername" field.
     * @type {string}
     * @memberof EntCustomer
     */
    customername?: string;
    /**
     * Identificationnumber holds the value of the "Identificationnumber" field.
     * @type {string}
     * @memberof EntCustomer
     */
    identificationnumber?: string;
    /**
     * Phonenumber holds the value of the "Phonenumber" field.
     * @type {string}
     * @memberof EntCustomer
     */
    phonenumber?: string;
    /**
     * 
     * @type {EntCustomerEdges}
     * @memberof EntCustomer
     */
    edges?: EntCustomerEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCustomer
     */
    id?: number;
}

export function EntCustomerFromJSON(json: any): EntCustomer {
    return EntCustomerFromJSONTyped(json, false);
}

export function EntCustomerFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCustomer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'address': !exists(json, 'Address') ? undefined : json['Address'],
        'customername': !exists(json, 'Customername') ? undefined : json['Customername'],
        'identificationnumber': !exists(json, 'Identificationnumber') ? undefined : json['Identificationnumber'],
        'phonenumber': !exists(json, 'Phonenumber') ? undefined : json['Phonenumber'],
        'edges': !exists(json, 'edges') ? undefined : EntCustomerEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCustomerToJSON(value?: EntCustomer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Address': value.address,
        'Customername': value.customername,
        'Identificationnumber': value.identificationnumber,
        'Phonenumber': value.phonenumber,
        'edges': EntCustomerEdgesToJSON(value.edges),
        'id': value.id,
    };
}


