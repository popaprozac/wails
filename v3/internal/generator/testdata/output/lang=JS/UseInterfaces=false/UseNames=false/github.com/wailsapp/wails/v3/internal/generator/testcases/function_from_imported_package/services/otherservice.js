// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * OtherService is a struct
 * that does things
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./internal.js";

/**
 * Yay does this and that
 * @returns {Promise<$models.Address | null> & { cancel(): void }}
 */
export function Yay() {
    let $resultPromise = /** @type {any} */($Call.ByID(2007737399));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType1($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = $models.Address.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
