<template>
  <UContainer>
    <h1 class="text-center">How to add devices</h1>
    <UDivider />
    <section id="master" class="w-fit p-4">
      <h2>Master</h2>
      <UDivider />
      <p>
        Bagian ini digunakan apabila pengguna belum terdaftar sehingga belum
        memiliki kubernetes cluster.
      </p>
      <ol>
        <li>Persiapkan sebuah server ataupun device yang memiliki</li>
        <li>Jalankan sebuah terminal lalu execute script ini.</li>
        <pre><code>curl -sfL https://get.k3s.io | sh -</code></pre>
        <p>
          Untuk detail lebih lengkapnya dapat dilihat
          <a href="https://k3s.io/"> k3s official page </a>
        </p>
        <li>
          Setelah mengeksekusi perintah, node master akan secara otomatis
          berjalan di background
        </li>
        <li>
          Pastikan server ataupun device dapat diakses secara publik oleh user
          ataupun perangkat lainnya.
        </li>
      </ol>
    </section>
    <section id="node" class="w-fit p-4">
      <h2>Node</h2>
      <UDivider />
      <p>
        Bagian ini digunakan apabila pengguna sudah memiliki kubernetes cluster
        yang terhubung.
      </p>
      <ol>
        <li>Persiapkan sebuah device untuk didaftarkan</li>
        <li>Jalankan perintah ini pada terminal</li>
        <pre><code>ping IP_PUBLIK_SERVER</code></pre>
        <p>
          Jika perintah memberikan reply, maka server sudah dapat diakses secara
          publik.
        </p>
        <li>
          Daftarkan sebagai node pada kubernetes cluster dengan menggunakan
          perintah ini
        </li>
        <pre><code>curl -sfL https://get.k3s.io | K3S_URL=https://ipserver:6443 K3S_TOKEN=mynodetoken K3S_NODE_NAME=mynodename sh -</code></pre>
        <div>
          dengan penjelasan:
          <ul>
            <li>K3S_URL: Alamat server</li>
            <li>K3S_TOKEN: Token yang didapat dari server</li>
            <li>K3S_NODE_NAME: Nama Node</li>
          </ul>
          Untuk detail lebih lengkapnya dapat dilihat
          <a href="https://k3s.io/"> k3s official page </a>
        </div>
        <li>
          Setelah mengeksekusi perintah, node akan terdaftar pada cluster secara
          otomatis
        </li>
      </ol>
    </section>
  </UContainer>
</template>

<style scoped>
ol {
  list-style: decimal;
}

ul {
  list-style: circle;
}

li {
  margin-left: 1rem;
  padding: 0.25rem;
}

a {
  cursor: pointer;
  color: gray;
  transition: all 0.1s ease-in-out;
  text-decoration: underline;
}

a:hover {
  color: #353845;
}

pre {
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  background-color: #353845;
}
</style>
