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
    EntFixEdges,
    EntFixEdgesFromJSON,
    EntFixEdgesFromJSONTyped,
    EntFixEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFix
 */
export interface EntFix {
    /**
     * Date holds the value of the "Date" field.
     * @type {string}
     * @memberof EntFix
     */
    date?: string;
    /**
     * Problemtype holds the value of the "Problemtype" field.
     * @type {string}
     * @memberof EntFix
     */
    problemtype?: string;
    /**
     * Productnumber holds the value of the "Productnumber" field.
     * @type {string}
     * @memberof EntFix
     */
    productnumber?: string;
    /**
     * Queue holds the value of the "Queue" field.
     * @type {string}
     * @memberof EntFix
     */
    queue?: string;
    /**
     * 
     * @type {EntFixEdges}
     * @memberof EntFix
     */
    edges?: EntFixEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntFix
     */
    id?: number;
}

export function EntFixFromJSON(json: any): EntFix {
    return EntFixFromJSONTyped(json, false);
}

export function EntFixFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFix {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'date': !exists(json, 'Date') ? undefined : json['Date'],
        'problemtype': !exists(json, 'Problemtype') ? undefined : json['Problemtype'],
        'productnumber': !exists(json, 'Productnumber') ? undefined : json['Productnumber'],
        'queue': !exists(json, 'Queue') ? undefined : json['Queue'],
        'edges': !exists(json, 'edges') ? undefined : EntFixEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntFixToJSON(value?: EntFix | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Date': value.date,
        'Problemtype': value.problemtype,
        'Productnumber': value.productnumber,
        'Queue': value.queue,
        'edges': EntFixEdgesToJSON(value.edges),
        'id': value.id,
    };
}


