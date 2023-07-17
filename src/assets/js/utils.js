// 通过时间戳获取格式化的时间
export function getDate(timestamp, needTime = false) {
    // 创建一个新的 Date 对象，并传入时间戳作为参数
    const date = new Date(timestamp * 1000);

    // 需要时间
    if (needTime) return date.toLocaleString()

    // 提取年、月、日
    const year = date.getFullYear();
    const month = date.getMonth() + 1; // 月份是从0开始的，需要加1
    const day = date.getDate();

    // 构建日期字符串
    const dateString = `${year}-${month}-${day}`;
    return dateString
}

// 获取该文章字数，返回字符串
export function getWordCount(content) {
    let len = content.length
    if (len < 1000) return String(len)
    else if (len < 10000) {
        // 1k - 1w
        return (len / 1000).toFixed(1) + ' k'
    }
    return (len / 10000).toFixed(1) + ' w'
}