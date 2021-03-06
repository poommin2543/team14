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
    EntPaymentTypeEdges,
    EntPaymentTypeEdgesFromJSON,
    EntPaymentTypeEdgesFromJSONTyped,
    EntPaymentTypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPaymentType
 */
export interface EntPaymentType {
    /**
     * Typename holds the value of the "Typename" field.
     * @type {string}
     * @memberof EntPaymentType
     */
    typename?: string;
    /**
     * 
     * @type {EntPaymentTypeEdges}
     * @memberof EntPaymentType
     */
    edges?: EntPaymentTypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPaymentType
     */
    id?: number;
}

export function EntPaymentTypeFromJSON(json: any): EntPaymentType {
    return EntPaymentTypeFromJSONTyped(json, false);
}

export function EntPaymentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPaymentType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'typename': !exists(json, 'Typename') ? undefined : json['Typename'],
        'edges': !exists(json, 'edges') ? undefined : EntPaymentTypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPaymentTypeToJSON(value?: EntPaymentType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Typename': value.typename,
        'edges': EntPaymentTypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}


