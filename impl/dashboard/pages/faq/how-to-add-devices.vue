<template>
  <UContainer>
    <h1 class="text-center">How to add devices</h1>
    <UDivider />
    <section id="master" class="w-fit p-4">
      <h2>Master</h2>
      <UDivider />
      <p>
        This section is used if the user is not yet registered and therefore
        does not have a Kubernetes cluster.
      </p>
      <ol>
        <li>
          Prepare a server or a device that capabale to connect to internet
        </li>
        <li>Run a terminal and execute this script.</li>
        <pre><code>curl -sfL https://get.k3s.io | sh -</code></pre>
        <p>
          For more detailed information, please refer to the
          <a href="https://k3s.io/"> k3s official page </a>
        </p>
        <li>
          After executing the command, the master node will automatically run in
          the background.
        </li>
        <li>
          Ensure the server or device can be publicly accessed by users or other
          devices.
        </li>
      </ol>
    </section>
    <section id="node" class="w-fit p-4">
      <h2>Node</h2>
      <UDivider />
      <p>
        This section is used if the user already has a connected Kubernetes
        cluster.
      </p>
      <ol>
        <li>Prepare a device to be registered</li>
        <li>Run this command in the terminal</li>
        <pre><code>ping PUBLIC_SERVER_IP</code></pre>
        <p>
          If the command returns a reply, the server is already publicly
          accessible.
        </p>
        <li>Register as a node in the Kubernetes cluster using this command</li>
        <pre><code>curl -sfL https://get.k3s.io | K3S_URL=https://ipserver:6443 K3S_TOKEN=mynodetoken K3S_NODE_NAME=mynodename sh -</code></pre>
        <div>
          with the explanation:
          <ul>
            <li>K3S_URL: Server address</li>
            <li>K3S_TOKEN: Token obtained from the server</li>
            <li>K3S_NODE_NAME: Node name</li>
          </ul>
          For more detailed information, please refer to the
          <a href="https://k3s.io/"> k3s official page </a>
        </div>
        <li>
          After executing the command, the node will be automatically registered
          to the cluster.
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
