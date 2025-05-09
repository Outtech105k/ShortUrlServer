document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("urlForm");

    form.addEventListener("submit", async function (e) {
        e.preventDefault();

        const activeTab = document.querySelector(".nav-link.active").id;
        const expireValue = document.getElementById("expire_value").value.trim();
        const expireUnit = document.getElementById("expire_unit").value;
        const expire_in = expireValue ? `${expireValue}${expireUnit}` : null;

        const body = {
            base_url: document.getElementById("base_url").value,
            expire_in: expire_in
        };

        if (activeTab === "custom-tab") {
            const customId = document.getElementById("custom_id").value.trim();
            if (!customId) {
                alert("カスタムIDを入力してください。");
                return;
            }
            body.custom_id = customId;
        } else {
            body.use_uppercase = document.getElementById("use_uppercase").checked;
            body.use_lowercase = document.getElementById("use_lowercase").checked;
            body.use_numbers = document.getElementById("use_numbers").checked;
            body.id_length = parseInt(document.getElementById("id_length").value);
        }

        try {
            const res = await fetch("/api/shorturl/set", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            });

            const data = await res.json();

            if (res.ok) {
                document.getElementById("short_url").value = data.short_url;
                document.getElementById("jump_link").href = data.short_url;
                document.getElementById("result").style.display = "block";
            } else {
                alert("エラー: " + data.error);
            }
        } catch (err) {
            alert("リクエスト失敗: " + err.message);
        }
    });
});

function copyToClipboard() {
    const urlField = document.getElementById("short_url");
    urlField.select();
    document.execCommand("copy");
}
