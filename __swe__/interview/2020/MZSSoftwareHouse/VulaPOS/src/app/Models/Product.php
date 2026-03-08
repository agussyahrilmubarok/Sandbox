<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Casts\Attribute;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Product extends Model
{
    use HasFactory;

    protected $fillable = [
        'title',
        'description',
        'stock',
        'buy_price',
        'sell_price',
        'image',
        'barcode',
        'category_id'
    ];

    // Relationship with Category
    public function category(): BelongsTo
    {
        return $this->belongsTo(Category::class);
    }

    // Handle image in storage
    protected function image(): Attribute
    {
        return Attribute::make(
            get: fn ($value) => asset('/storage/products/' . $value),
        );
    }
}
