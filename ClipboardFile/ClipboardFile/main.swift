//
//  main.swift
//  ClipboardFile
//
//  Created by Sindre Michael Andreassen on 18/02/2024.
//

import Foundation
import Cocoa


let pasteboard = NSPasteboard.general
let items = pasteboard.pasteboardItems

if let items = items {
    for item in items {
        if let urlString = item.string(forType: .fileURL), let url = URL(string: urlString) {
            print(url.path)
        }
    }
}
