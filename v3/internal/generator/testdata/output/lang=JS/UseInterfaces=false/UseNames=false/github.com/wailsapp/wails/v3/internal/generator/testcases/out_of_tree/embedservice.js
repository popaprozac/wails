// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * EmbedService is tricky.
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as nobindingshere$0 from "../no_bindings_here/models.js";

/**
 * LikeThisOne is an example method that does nothing.
 * @returns {Promise<[nobindingshere$0.Person, nobindingshere$0.HowDifferent<boolean>, nobindingshere$0.PrivatePerson]> & { cancel(): void }}
 */
export function LikeThisOne() {
    let $resultPromise = /** @type {any} */($Call.ByID(2590614085));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        $result[0] = $$createType0($result[0]);
        $result[1] = $$createType1($result[1]);
        $result[2] = $$createType2($result[2]);
        return $result;
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * LikeThisOtherOne does nothing as well, but is different.
 * @returns {Promise<void> & { cancel(): void }}
 */
export function LikeThisOtherOne() {
    let $resultPromise = /** @type {any} */($Call.ByID(773650321));
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = nobindingshere$0.Person.createFrom;
const $$createType1 = nobindingshere$0.HowDifferent.createFrom($Create.Any);
const $$createType2 = nobindingshere$0.PrivatePerson.createFrom;
