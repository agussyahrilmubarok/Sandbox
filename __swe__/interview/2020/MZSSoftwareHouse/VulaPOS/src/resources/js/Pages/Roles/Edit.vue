<script setup>
import { defineProps, } from 'vue';
import { Head, useForm } from '@inertiajs/vue3';
import DashboardLayout from '@/Layouts/DashboardLayout.vue';
import Swal from 'sweetalert2';

const props = defineProps({
    errors: {
        type: Object,
    },
    role: {
        type: Object,
    },
    permissions: {
        type: Array
    }
});

const form = useForm({
    name: props.role.name,
    permissions: props.role.permissions.map(obj => obj.name),
});

const onSubmit = () => {
    form.put(`/dashboard/roles/${props.role.id}`, form, {
        onSuccess: () => {
            Swal.fire({
                title: 'Success!',
                text: 'Role updated successfully.',
                icon: 'success',
                showConfirmButton: false,
                timer: 2000
            });
        },
    });
}

</script>

<template>
    <DashboardLayout>

        <Head title="Edit Role" />

        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card border-0 rounded-3 shadow border-top-purple">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-shield-alt"></i> EDIT ROLE</span>
                                </div>
                                <div class="card-body">

                                    <form @submit.prevent="onSubmit">

                                        <div class="mb-3">
                                            <label class="fw-bold">Role Name</label>
                                            <input class="form-control" v-model="form.name"
                                                :class="{ 'is-invalid': errors.name }" type="text"
                                                placeholder="Role Name">

                                            <div v-if="errors.name" class="alert alert-danger">
                                                {{ errors.name }}
                                            </div>
                                        </div>

                                        <hr>
                                        <div class="mb-3">
                                            <label class="fw-bold">Permissions</label>
                                            <br>
                                            <div class="form-check form-check-inline"
                                                v-for="(permission, index) in permissions" :key="index">
                                                <input class="form-check-input" type="checkbox"
                                                    v-model="form.permissions" :value="permission.name"
                                                    :id="`check-${permission.id}`">
                                                <label class="form-check-label" :for="`check-${permission.id}`">{{
                                                    permission.name }}</label>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-12">
                                                <button class="btn btn-primary shadow-sm rounded-sm"
                                                    type="submit">UPDATE</button>
                                                <button class="btn btn-warning shadow-sm rounded-sm ms-3"
                                                    type="reset">RESET</button>
                                            </div>
                                        </div>
                                    </form>

                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>

    </DashboardLayout>
</template>
