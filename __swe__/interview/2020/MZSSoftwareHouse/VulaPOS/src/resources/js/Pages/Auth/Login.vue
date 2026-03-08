<script setup>
import { Head, Link, useForm } from '@inertiajs/vue3'
import AuthLayout from '@/Layouts/AuthLayout.vue'
import ApplicationLogo from '@/Components/ApplicationLogo.vue';
import InputError from '@/Components/InputError.vue';

defineProps({
    errors: {
        type: Object,
    },
    status: {
        type: String,
    },
});

const form = useForm({
    email: '',
    password: '',
});

const onSubmit = () => {
    form.post(route('login'), {
        onFinish: () => form.reset('password'),
    });
}
</script>

<template>
    <AuthLayout>

        <Head title="Login" />

        <div class="d-flex min-vh-100 flex-column justify-content-center align-items-center">
            <div class="col-md-8 col-lg-5">
                <div class="fade-in">

                    <div class="text-center my-4">
                        <a href="" class="text-dark text-decoration-none">
                            <ApplicationLogo style="width: 80px; height: 80px;" />
                            <h3 class="mt-2 font-weight-bold">Vula POS</h3>
                        </a>
                    </div>

                    <div class="card border-top-dark border-0 shadow-sm rounded-3">
                        <div class="card-body">

                            <div class="text-start">
                                <p class="text-muted text-center">Please use your account to login!</p>
                            </div>
                            <hr>
                            <div v-if="status" class="alert alert-success mt-2">
                                {{ status }}
                            </div>

                            <form @submit.prevent="onSubmit">
                                <div class="input-group mb-3">
                                    <span class="input-group-text" id="basic-addon1">
                                        <i class="fa fa-envelope"></i>
                                    </span>
                                    <input type="email" v-model="form.email" id="email" class="form-control"
                                        placeholder="Your email address" aria-label="Email"
                                        aria-describedby="basic-addon1" />
                                </div>
                                <InputError class="mt-2" :message="form.errors.email" />

                                <div class="input-group mb-3">
                                    <span class="input-group-text" id="basic-addon1">
                                        <i class="fa fa-lock"></i>
                                    </span>
                                    <input type="password" v-model="form.password" id="password" class="form-control"
                                        placeholder="Your password" aria-label="Password"
                                        aria-describedby="basic-addon1" />
                                </div>
                                <InputError :message="form.errors.password" />

                                <div class="row">
                                    <div class="col-12 mb-3 text-end">
                                        <Link :href="route('password.request')">Forgot Password?</Link>
                                    </div>
                                    <div class="col-12">
                                        <button class="btn btn-dark shadow-sm rounded-sm px-4 w-100" type="submit">
                                            Login
                                        </button>
                                    </div>
                                </div>
                            </form>

                        </div>
                    </div>

                </div>
            </div>
        </div>
    </AuthLayout>
</template>
