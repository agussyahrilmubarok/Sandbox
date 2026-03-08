<?php

namespace Database\Factories;

use App\Models\Product;
use App\Models\Transaction;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\TransactionDetail>
 */
class TransactionDetailFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'qty' => $this->faker->numberBetween(1, 100),
            'price' => $this->faker->numberBetween(100, 10000),
            'transaction_id' => Transaction::factory(), // Create a Transaction record for transaction_id
            'product_id' => Product::factory(), // Create a Product record for product_id
        ];
    }
}
