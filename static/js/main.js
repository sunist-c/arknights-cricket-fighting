// 全局配置
const API_BASE_URL = '/v1';

// 工具函数
const showToast = (message, type = 'success') => {
    const toast = document.createElement('div');
    toast.className = `toast align-items-center text-white bg-${type} border-0 position-fixed top-0 end-0 m-3`;
    toast.setAttribute('role', 'alert');
    toast.setAttribute('aria-live', 'assertive');
    toast.setAttribute('aria-atomic', 'true');
    
    toast.innerHTML = `
        <div class="d-flex">
            <div class="toast-body">
                ${message}
            </div>
            <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast"></button>
        </div>
    `;
    
    document.body.appendChild(toast);
    const bsToast = new bootstrap.Toast(toast);
    bsToast.show();
    
    toast.addEventListener('hidden.bs.toast', () => {
        toast.remove();
    });
};

// 表单提交处理
const handleFormSubmit = async (formId, endpoint, successCallback) => {
    const form = document.getElementById(formId);
    if (!form) return;

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());

        try {
            const response = await fetch(`${API_BASE_URL}${endpoint}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            const result = await response.json();
            
            if (response.ok) {
                showToast(result.message || '操作成功');
                if (successCallback) successCallback(result);
            } else {
                showToast(result.message || '操作失败', 'danger');
            }
        } catch (error) {
            showToast('网络错误，请稍后重试', 'danger');
            console.error('Error:', error);
        }
    });
};

// 页面加载完成后执行
document.addEventListener('DOMContentLoaded', () => {
    // 注册表单处理
    handleFormSubmit('registerForm', '/common/register', (result) => {
        // 延迟跳转，让用户看到成功提示
        setTimeout(() => {
            window.location.href = '/login';
        }, 1500);
    });

    // 地址表单处理
    handleFormSubmit('addressForm', '/common/address', (result) => {
        // 刷新页面或更新地址信息
        window.location.reload();
    });

    // 报名表单处理
    handleFormSubmit('participantForm', '/cricket-fighting-6th/participant', (result) => {
        // 刷新页面或更新报名状态
        window.location.reload();
    });

    // 活动规则接受处理
    handleFormSubmit('licenseForm', '/cricket-fighting-6th/license-accept', (result) => {
        // 刷新页面或更新状态
        window.location.reload();
    });
}); 