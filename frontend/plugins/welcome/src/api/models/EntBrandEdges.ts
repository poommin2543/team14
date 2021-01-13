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
    EntProduct,
    EntProductFromJSON,
    EntProductFromJSONTyped,
    EntProductToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBrandEdges
 */
export interface EntBrandEdges {
    /**
     * Fix holds the value of the fix edge.
     * @type {Array<EntFix>}
     * @memberof EntBrandEdges
     */
    fix?: Array<EntFix>;
    /**
     * Product holds the value of the product edge.
     * @type {Array<EntProduct>}
     * @memberof EntBrandEdges
     */
    product?: Array<EntProduct>;
}

export function EntBrandEdgesFromJSON(json: any): EntBrandEdges {
    return EntBrandEdgesFromJSONTyped(json, false);
}

export function EntBrandEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBrandEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fix': !exists(json, 'fix') ? undefined : ((json['fix'] as Array<any>).map(EntFixFromJSON)),
        'product': !exists(json, 'product') ? undefined : ((json['product'] as Array<any>).map(EntProductFromJSON)),
    };
}

export function EntBrandEdgesToJSON(value?: EntBrandEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fix': value.fix === undefined ? undefined : ((value.fix as Array<any>).map(EntFixToJSON)),
        'product': value.product === undefined ? undefined : ((value.product as Array<any>).map(EntProductToJSON)),
    };
}


