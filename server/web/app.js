let dict = {}
let item = {}

$('document').ready(() => {

    fetch('/serialize').then(rsp => {
        return rsp.json()
    }).then(rsp => {
        dict = rsp.dict
        item = rsp
        delete item.dict

        console.log(dict, item)
        return item
    }).then(item => {
        renderItem(item, $('#canvas'))
    })
})

function renderItem(item, $canvas) {
    if (item.type.startsWith('*')) {
        renderItem(dict[item.value], $canvas)
        return
    }

    const $itemDom = $('<div><div>')
    $canvas.html($itemDom)
    switch (item.type) {
        case "int":
        case "int8":
        case "int16":
        case "int32":
        case "int64":
        case "uint":
        case "uint8":
        case "uint16":
        case "uint32":
        case "uint64":
        case "float32":
        case "float64":
            $itemDom.append($(`<div>${item.value}</div>`))
            break;
        default:
            $ul = $('<ul></ul>')
            $itemDom.append($ul)
            for (let key in item.value) {
                const $entryDom = $('<li></li>')
                $ul.append($entryDom)
                const $entryDomKey = $(`<div>${key}</div>`)
                const $entryDomDiv = $('<div></div>')
                $entryDom.append($entryDomKey)
                $entryDom.append($entryDomDiv)
                $entryDomKey.click(() => {
                    renderItem(item.value[key], $entryDomDiv)
                })
            }
            break;
    }
}