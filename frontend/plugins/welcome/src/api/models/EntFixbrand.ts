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
    EntFixbrandEdges,
    EntFixbrandEdgesFromJSON,
    EntFixbrandEdgesFromJSONTyped,
    EntFixbrandEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFixbrand
 */
export interface EntFixbrand {
    /**
     * 
     * @type {EntFixbrandEdges}
     * @memberof EntFixbrand
     */
    edges?: EntFixbrandEdges;
    /**
     * Fixbrandname holds the value of the "fixbrandname" field.
     * @type {string}
     * @memberof EntFixbrand
     */
    fixbrandname?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntFixbrand
     */
    id?: number;
}

export function EntFixbrandFromJSON(json: any): EntFixbrand {
    return EntFixbrandFromJSONTyped(json, false);
}

export function EntFixbrandFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFixbrand {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntFixbrandEdgesFromJSON(json['edges']),
        'fixbrandname': !exists(json, 'fixbrandname') ? undefined : json['fixbrandname'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntFixbrandToJSON(value?: EntFixbrand | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntFixbrandEdgesToJSON(value.edges),
        'fixbrandname': value.fixbrandname,
        'id': value.id,
    };
}


