<script lang="ts">
  import axios from 'axios';
  import { api } from './api';
  import Form from './components/Form.svelte';

  const signup = async (email: string, password: string) => {
    const { data } = await api.post('/auth/signup', { email, password });

    console.log(data);
  };

  const signin = async (email: string, password: string) => {
    const { data } = await api.post('/auth/signin', { email, password });

    console.log(data);
  };
</script>

<main>
  <Form onSubmit={signup} />
  <Form onSubmit={signin} actionTitle="Sign in" />
</main>
