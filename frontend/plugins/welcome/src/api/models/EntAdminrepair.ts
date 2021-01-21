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
    EntAdminrepairEdges,
    EntAdminrepairEdgesFromJSON,
    EntAdminrepairEdgesFromJSONTyped,
    EntAdminrepairEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAdminrepair
 */
export interface EntAdminrepair {
    /**
     * 
     * @type {EntAdminrepairEdges}
     * @memberof EntAdminrepair
     */
    edges?: EntAdminrepairEdges;
    /**
     * Equipmentdamate holds the value of the "equipmentdamate" field.
     * @type {string}
     * @memberof EntAdminrepair
     */
    equipmentdamate?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntAdminrepair
     */
    id?: number;
    /**
     * Numberrepair holds the value of the "numberrepair" field.
     * @type {string}
     * @memberof EntAdminrepair
     */
    numberrepair?: string;
    /**
     * Repairinformation holds the value of the "repairinformation" field.
     * @type {string}
     * @memberof EntAdminrepair
     */
    repairinformation?: string;
}

export function EntAdminrepairFromJSON(json: any): EntAdminrepair {
    return EntAdminrepairFromJSONTyped(json, false);
}

export function EntAdminrepairFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAdminrepair {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntAdminrepairEdgesFromJSON(json['edges']),
        'equipmentdamate': !exists(json, 'equipmentdamate') ? undefined : json['equipmentdamate'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'numberrepair': !exists(json, 'numberrepair') ? undefined : json['numberrepair'],
        'repairinformation': !exists(json, 'repairinformation') ? undefined : json['repairinformation'],
    };
}

export function EntAdminrepairToJSON(value?: EntAdminrepair | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntAdminrepairEdgesToJSON(value.edges),
        'equipmentdamate': value.equipmentdamate,
        'id': value.id,
        'numberrepair': value.numberrepair,
        'repairinformation': value.repairinformation,
    };
}


