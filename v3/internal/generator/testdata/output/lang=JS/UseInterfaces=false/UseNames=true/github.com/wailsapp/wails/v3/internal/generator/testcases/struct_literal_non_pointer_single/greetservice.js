// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * GreetService is great
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./internal.js";

/**
 * @param {number[]} $in
 * @returns {Promise<void> & { cancel(): void }}
 */
export function ArrayInt($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.ArrayInt", $in));
    return $resultPromise;
}

/**
 * @param {boolean} $in
 * @returns {Promise<boolean> & { cancel(): void }}
 */
export function BoolInBoolOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.BoolInBoolOut", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function Float32InFloat32Out($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Float32InFloat32Out", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function Float64InFloat64Out($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Float64InFloat64Out", $in));
    return $resultPromise;
}

/**
 * Greet someone
 * @param {string} name
 * @returns {Promise<string> & { cancel(): void }}
 */
export function Greet(name) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Greet", name));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function Int16InIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int16InIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function Int16PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int16PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function Int32InIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int32InIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function Int32PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int32PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function Int64InIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int64InIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function Int64PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int64PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function Int8InIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int8InIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function Int8PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.Int8PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function IntInIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.IntInIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function IntPointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.IntPointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function IntPointerInputNamedOutputs($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.IntPointerInputNamedOutputs", $in));
    return $resultPromise;
}

/**
 * @param {{ [_: `${number}`]: number }} $in
 * @returns {Promise<void> & { cancel(): void }}
 */
export function MapIntInt($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.MapIntInt", $in));
    return $resultPromise;
}

/**
 * @param {{ [_: string]: number }} $in
 * @returns {Promise<void> & { cancel(): void }}
 */
export function MapIntPointerInt($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.MapIntPointerInt", $in));
    return $resultPromise;
}

/**
 * @param {{ [_: `${number}`]: number[] }} $in
 * @returns {Promise<void> & { cancel(): void }}
 */
export function MapIntSliceInt($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.MapIntSliceInt", $in));
    return $resultPromise;
}

/**
 * @param {{ [_: `${number}`]: number[] }} $in
 * @returns {Promise<{ [_: `${number}`]: number[] }> & { cancel(): void }}
 */
export function MapIntSliceIntInMapIntSliceIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.MapIntSliceIntInMapIntSliceIntOut", $in));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType1($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * @returns {Promise<string> & { cancel(): void }}
 */
export function NoInputsStringOut() {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.NoInputsStringOut"));
    return $resultPromise;
}

/**
 * @param {boolean | null} $in
 * @returns {Promise<boolean | null> & { cancel(): void }}
 */
export function PointerBoolInBoolOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.PointerBoolInBoolOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function PointerFloat32InFloat32Out($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.PointerFloat32InFloat32Out", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function PointerFloat64InFloat64Out($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.PointerFloat64InFloat64Out", $in));
    return $resultPromise;
}

/**
 * @param {{ [_: `${number}`]: number } | null} $in
 * @returns {Promise<void> & { cancel(): void }}
 */
export function PointerMapIntInt($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.PointerMapIntInt", $in));
    return $resultPromise;
}

/**
 * @param {string | null} $in
 * @returns {Promise<string | null> & { cancel(): void }}
 */
export function PointerStringInStringOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.PointerStringInStringOut", $in));
    return $resultPromise;
}

/**
 * @param {string[]} $in
 * @returns {Promise<string[]> & { cancel(): void }}
 */
export function StringArrayInputNamedOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StringArrayInputNamedOutput", $in));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType2($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * @param {string[]} $in
 * @returns {Promise<string[]> & { cancel(): void }}
 */
export function StringArrayInputNamedOutputs($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StringArrayInputNamedOutputs", $in));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType2($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * @param {string[]} $in
 * @returns {Promise<string[]> & { cancel(): void }}
 */
export function StringArrayInputStringArrayOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StringArrayInputStringArrayOut", $in));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType2($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * @param {string[]} $in
 * @returns {Promise<string> & { cancel(): void }}
 */
export function StringArrayInputStringOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StringArrayInputStringOut", $in));
    return $resultPromise;
}

/**
 * @param {$models.Person} $in
 * @returns {Promise<$models.Person> & { cancel(): void }}
 */
export function StructInputStructOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StructInputStructOutput", $in));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType3($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * @param {$models.Person | null} $in
 * @returns {Promise<void> & { cancel(): void }}
 */
export function StructPointerInputErrorOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StructPointerInputErrorOutput", $in));
    return $resultPromise;
}

/**
 * @param {$models.Person | null} $in
 * @returns {Promise<$models.Person | null> & { cancel(): void }}
 */
export function StructPointerInputStructPointerOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.StructPointerInputStructPointerOutput", $in));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType4($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function UInt16InUIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt16InUIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function UInt16PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt16PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function UInt32InUIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt32InUIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function UInt32PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt32PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function UInt64InUIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt64InUIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function UInt64PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt64PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function UInt8InUIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt8InUIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function UInt8PointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UInt8PointerInAndOutput", $in));
    return $resultPromise;
}

/**
 * @param {number} $in
 * @returns {Promise<number> & { cancel(): void }}
 */
export function UIntInUIntOut($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UIntInUIntOut", $in));
    return $resultPromise;
}

/**
 * @param {number | null} $in
 * @returns {Promise<number | null> & { cancel(): void }}
 */
export function UIntPointerInAndOutput($in) {
    let $resultPromise = /** @type {any} */($Call.ByName("main.GreetService.UIntPointerInAndOutput", $in));
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = $Create.Array($Create.Any);
const $$createType1 = $Create.Map($Create.Any, $$createType0);
const $$createType2 = $Create.Array($Create.Any);
const $$createType3 = $models.Person.createFrom;
const $$createType4 = $Create.Nullable($$createType3);
