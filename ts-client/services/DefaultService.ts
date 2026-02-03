/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ErrorModel } from '../models/ErrorModel';
import type { LockOutputBody } from '../models/LockOutputBody';
import type { StatusOutputBody } from '../models/StatusOutputBody';
import type { UnlockOutputBody } from '../models/UnlockOutputBody';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * Post lock
     * @returns LockOutputBody OK
     * @returns ErrorModel Error
     * @throws ApiError
     */
    public static postLock(): CancelablePromise<LockOutputBody | ErrorModel> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/lock',
        });
    }
    /**
     * Get status
     * @returns StatusOutputBody OK
     * @returns ErrorModel Error
     * @throws ApiError
     */
    public static getStatus(): CancelablePromise<StatusOutputBody | ErrorModel> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/status',
        });
    }
    /**
     * Post unlock
     * @returns UnlockOutputBody OK
     * @returns ErrorModel Error
     * @throws ApiError
     */
    public static postUnlock(): CancelablePromise<UnlockOutputBody | ErrorModel> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/unlock',
        });
    }
}
