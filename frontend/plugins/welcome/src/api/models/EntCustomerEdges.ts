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
    EntFix,
    EntFixFromJSON,
    EntFixFromJSONTyped,
    EntFixToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntPersonal,
    EntPersonalFromJSON,
    EntPersonalFromJSONTyped,
    EntPersonalToJSON,
    EntTitle,
    EntTitleFromJSON,
    EntTitleFromJSONTyped,
    EntTitleToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCustomerEdges
 */
export interface EntCustomerEdges {
    /**
     * Fix holds the value of the fix edge.
     * @type {Array<EntFix>}
     * @memberof EntCustomerEdges
     */
    fix?: Array<EntFix>;
    /**
     * 
     * @type {EntGender}
     * @memberof EntCustomerEdges
     */
    gender?: EntGender;
    /**
     * 
     * @type {EntPersonal}
     * @memberof EntCustomerEdges
     */
    personal?: EntPersonal;
    /**
     * 
     * @type {EntTitle}
     * @memberof EntCustomerEdges
     */
    title?: EntTitle;
}

export function EntCustomerEdgesFromJSON(json: any): EntCustomerEdges {
    return EntCustomerEdgesFromJSONTyped(json, false);
}

export function EntCustomerEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCustomerEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fix': !exists(json, 'Fix') ? undefined : ((json['Fix'] as Array<any>).map(EntFixFromJSON)),
        'gender': !exists(json, 'Gender') ? undefined : EntGenderFromJSON(json['Gender']),
        'personal': !exists(json, 'Personal') ? undefined : EntPersonalFromJSON(json['Personal']),
        'title': !exists(json, 'Title') ? undefined : EntTitleFromJSON(json['Title']),
    };
}

export function EntCustomerEdgesToJSON(value?: EntCustomerEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fix': value.fix === undefined ? undefined : ((value.fix as Array<any>).map(EntFixToJSON)),
        'gender': EntGenderToJSON(value.gender),
        'personal': EntPersonalToJSON(value.personal),
        'title': EntTitleToJSON(value.title),
    };
}


