function extractContent() {
    const jsonInput = document.getElementById('jsonInput').value;
    const output = document.getElementById('output');
    
    try {
        // 解析JSON
        const data = JSON.parse(jsonInput);
        
        // 检查是否存在body数组
        if (data.body && Array.isArray(data.body)) {
            // 提取所有content内容
            const contents = data.body.map(item => item.content);
            
            // 显示结果
            output.textContent = contents.join('\n');
        } else {
            output.textContent = '错误：未找到有效的body数组';
        }
    } catch (error) {
        output.textContent = '错误：无效的JSON格式\n' + error.message;
    }
}

function saveToFile() {
    const output = document.getElementById('output').textContent;
    if (!output) {
        alert('没有可保存的内容！');
        return;
    }

    const fileType = document.getElementById('fileType').value;
    const fileName = `extracted_content.${fileType}`;
    
    // 创建Blob对象
    const blob = new Blob([output], { type: 'text/plain;charset=utf-8' });
    
    // 创建下载链接
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = fileName;
    
    // 添加到文档中并触发点击
    document.body.appendChild(link);
    link.click();
    
    // 清理
    document.body.removeChild(link);
    URL.revokeObjectURL(link.href);
} 