<template>
    <div class="p-6 bg-gray-100 min-h-screen flex items-center justify-center">
        <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
            <div class="mb-6">
                <h4 class="text-2xl font-semibold text-gray-700">Add Transactions</h4>
            </div>

            <ul class="bg-red-600" v-if="err != null">
                <li class="text-black p-1">{{ err }}</li>
            </ul>

            <div class="space-y-4">
                <div>
                    <label for="type" class="block text-gray-600 font-medium mb-1">Type</label>
                    <input id="type" type="text"
                        class="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                        v-model="model.transaction.type" />
                </div>
                <div>
                    <label for="ticker" class="block text-gray-600 font-medium mb-1">Ticker</label>
                    <input id="ticker" type="text"
                        class="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                        v-model="model.transaction.ticker" />
                </div>
                <div>
                    <label for="volume" class="block text-gray-600 font-medium mb-1">Volume</label>
                    <input id="volume" type="text"
                        class="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                        v-model="model.transaction.volume" />
                </div>
                <div>
                    <label for="price" class="block text-gray-600 font-medium mb-1">Price</label>
                    <input id="price" type="text"
                        class="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                        v-model="model.transaction.price" />
                </div>
                <div>
                    <label for="date" class="block text-gray-600 font-medium mb-1">Date</label>
                    <input id="date" type="date"
                        class="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                        v-model="model.transaction.date" />
                </div>
                <div>
                    <button
                        class="w-full bg-blue-500 text-white font-semibold py-2 rounded-lg hover:bg-blue-600 transition duration-300 mt-4"
                        @click="saveTransaction">
                        Save
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import axios from 'axios';
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router';

export default defineComponent({
    name: 'transactionCreate',
    setup() {
        const router = useRouter();
        return { router };
    },
    data() {
        return {
            err: null,
            model: {
                transaction: {
                    type: '',
                    ticker: '',
                    volume: '',
                    price: '',
                    date: '',
                }
            }
        }
    },
    methods: {
        saveTransaction() {
            axios.post('http://localhost:80/api/txn/add', {
                type: this.model.transaction.type,
                ticker: this.model.transaction.ticker,
                volume: Number(this.model.transaction.volume),
                price: Number(this.model.transaction.price),
                date: this.model.transaction.date,
            })
                .then(res => {
                    console.log(res)

                    this.model.transaction = {
                        type: '',
                        ticker: '',
                        volume: '',
                        price: '',
                        date: '',
                    }
                    this.$router.push('/');
                })
                .catch((error) => {
                    console.log(error);
                    if (error.response && (error.response.status === 400 || error.response.status === 500)) {
                        console.log(error.response.data.error);
                        this.err = error.response.data.error;
                    }
                    if (error.response?.data?.error) {
                        console.log(error.response.data.error);
                        this.err = error.response.data.error;
                    } else {
                        console.error('Unexpected error:', error);
                    }
                })
        }
    }
})

</script>