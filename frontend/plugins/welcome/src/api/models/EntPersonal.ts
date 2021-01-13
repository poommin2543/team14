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
    EntPersonalEdges,
    EntPersonalEdgesFromJSON,
    EntPersonalEdgesFromJSONTyped,
    EntPersonalEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPersonal
 */
export interface EntPersonal {
    /**
     * Email holds the value of the "Email" field.
     * @type {string}
     * @memberof EntPersonal
     */
    email?: string;
    /**
     * Password holds the value of the "Password" field.
     * @type {string}
     * @memberof EntPersonal
     */
    password?: string;
    /**
     * Personalname holds the value of the "Personalname" field.
     * @type {string}
     * @memberof EntPersonal
     */
    personalname?: string;
    /**
     * 
     * @type {EntPersonalEdges}
     * @memberof EntPersonal
     */
    edges?: EntPersonalEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPersonal
     */
    id?: number;
}

export function EntPersonalFromJSON(json: any): EntPersonal {
    return EntPersonalFromJSONTyped(json, false);
}

export function EntPersonalFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPersonal {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'Email') ? undefined : json['Email'],
        'password': !exists(json, 'Password') ? undefined : json['Password'],
        'personalname': !exists(json, 'Personalname') ? undefined : json['Personalname'],
        'edges': !exists(json, 'edges') ? undefined : EntPersonalEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPersonalToJSON(value?: EntPersonal | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Email': value.email,
        'Password': value.password,
        'Personalname': value.personalname,
        'edges': EntPersonalEdgesToJSON(value.edges),
        'id': value.id,
    };
}


