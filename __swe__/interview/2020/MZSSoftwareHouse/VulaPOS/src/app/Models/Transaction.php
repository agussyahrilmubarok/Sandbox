<?php

namespace App\Models;

use Carbon\Carbon;
use Illuminate\Database\Eloquent\Casts\Attribute;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Transaction extends Model
{
    use HasFactory;

    protected $fillable = [
        'invoice',
        'cash',
        'change',
        'discount',
        'grand_total',
        'cashier_id',
        'customer_id'
    ];

    // Relationship with Transaction Details
    public function details(): HasMany
    {
        return $this->hasMany(TransactionDetail::class);
    }

    // Relationship with Cashier
    public function cashier(): BelongsTo
    {
        return $this->belongsTo(User::class, 'cashier_id');
    }

    // Relationship with Customer
    public function customer(): BelongsTo
    {
        return $this->belongsTo(Customer::class);
    }

    public function profits(): HasMany
    {
        return $this->hasMany(Profit::class);
    }


    // Handle Datetime on CreatedAt
    protected function createdAt(): Attribute
    {
        return Attribute::make(
            get: fn ($value) => Carbon::parse($value)->format('d-M-Y H:i:s'),
        );
    }
}
