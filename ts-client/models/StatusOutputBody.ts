/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type StatusOutputBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    /**
     * Battery critical warning
     */
    battery_critical: boolean;
    /**
     * Battery percentage
     */
    battery_percent: number;
    /**
     * Current lock state
     */
    lock_state: string;
    /**
     * Nuki device state
     */
    nuki_state: string;
};

