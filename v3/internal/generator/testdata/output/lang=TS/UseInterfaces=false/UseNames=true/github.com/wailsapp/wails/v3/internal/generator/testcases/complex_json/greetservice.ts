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
 * Greet does XYZ
 */
export function Greet(person: $models.Person, emb: $models.Embedded1): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByName("main.GreetService.Greet", person, emb) as any;
    return $resultPromise;
}
