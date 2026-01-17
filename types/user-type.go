package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model `user` table
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	UserRole  string             `json:"user_role" bson:"user_role"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// address model `addresses` table
type Address struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Address1 string             `json:"address1" bson:"address1"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	City     string             `json:"city" bson:"city"`
	Country  string             `json:"country" bson:"country"`
}

// verification model `verification` table
type Verification struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	OTP       string             `json:"otp" bson:"otp"`
	Status    bool               `json:"status" bson:"status"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"` //MongoDB-র একটি ফিচার আছে TTL Index, যার মাধ্যমে আপনি এই ExpiresAt সময় পার হওয়ার সাথে সাথে রেকর্ডটি ডাটাবেজ থেকে অটোমেটিক ডিলিট করে দিতে পারেন।

	// 	৩. এই মডেলটি নিয়ে পরবর্তী ধাপ (Workflow)
	// আপনি যখন ওটিপি সিস্টেম বানাবেন, তখন এই লজিকটি ফলো করবেন:

	// Generate: একটি র‍্যান্ডম নম্বর তৈরি করবেন এবং সেটি ডাটাবেজে ExpiresAt সহ সেভ করবেন (Status থাকবে false)।

	// Send: ইউজারকে ই-মেইলে ওই OTP পাঠাবেন।

	// Verify: ইউজার যখন OTP দিবে, আপনি ডাটাবেজে চেক করবেন:

	// ই-মেইল এবং ওটিপি মিলছে কি না।

	// বর্তমান সময় ExpiresAt-এর চেয়ে কম কি না।

	// Update: সবকিছু ঠিক থাকলে Status কে true করে দিবেন।
}
