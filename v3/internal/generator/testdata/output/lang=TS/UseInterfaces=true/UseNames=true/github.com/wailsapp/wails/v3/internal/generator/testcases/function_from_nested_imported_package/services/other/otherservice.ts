// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * OtherService is a struct
 * that does things
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./internal.js";

/**
 * Yay does this and that
 */
export function Yay(): Promise<$models.Address | null> & { cancel(): void } {
    let $resultPromise = $Call.ByName("github.com/wailsapp/wails/v3/internal/generator/testcases/function_from_nested_imported_package/services/other.OtherService.Yay") as any;
    return $resultPromise;
}
