<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\DB;


class InitTables extends Migration
{
    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
    }

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('countries', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->primary('code');
            $table->string('name', 255);
            $table->string('code', 16);
            $table->string('currency_code', 8);
        });

        Schema::create('tiers', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->bigIncrements('id');

            $table->string('country_code', 16);
            $table->foreign('country_code')->references('code')->on('countries')->onDelete('cascade');

            $table->integer('level')->nullable(false);
            $table->string('name', 255);
        });

        Schema::create('tier_requirements', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->bigIncrements('id');

            $table->unsignedBigInteger('tier_id')->nullable(false);
            $table->foreign('tier_id')->references('id')->on('tiers')->onDelete('cascade');

            $table->string('name')->nullable(false);
            $table->string('form_index', 255);
            $table->text('options')->nullable(true);
        });

        Schema::create('user_requirements', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->bigIncrements('id');

            $table->unsignedBigInteger('tier_requirement_id')->nullable(false);
            $table->foreign('tier_requirement_id')->references('id')->on('tier_requirements')->onDelete('cascade');

            $table->string('user_id', 255)->nullable(false);
            $table->string('status', 36)->nullable(false);

            $table->timestamps();
        });

        Schema::create('user_requirement_values', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->bigIncrements('id');

            $table->unsignedBigInteger('user_requirement_id')->nullable(false);
            $table->foreign('user_requirement_id')->references('id')->on('user_requirements')->onDelete('cascade');

            $table->string('index', 255)->nullable(false);
            $table->text('value')->nullable(true);

            $table->timestamps();
        });

        Schema::create('tier_features', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->unsignedBigInteger('tier_id')->nullable(false);
            $table->foreign('tier_id')->references('id')->on('tiers')->onDelete('cascade');

            $table->string('index', 255)->nullable(false);
        });

        Schema::create('user_requests', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->bigIncrements('id');

            $table->unsignedBigInteger('tier_id')->nullable(false);
            $table->foreign('tier_id')->references('id')->on('tiers')->onDelete('cascade');

            $table->string('user_id', 255)->nullable(false);
            $table->string('status', 36)->nullable(false);

            $table->timestamps();
        });

        Schema::create('tier_limitations', function (Blueprint $table) {
            $table->charset = 'utf8';
            $table->collation = 'utf8_general_ci';

            $table->bigIncrements('id');

            $table->unsignedBigInteger('tier_id')->nullable(false);
            $table->foreign('tier_id')->references('id')->on('tiers')->onDelete('cascade');

             $table->decimal('value', 36, 18)->nullable(true);
            $table->string('index', 45)->nullable(false);
        });

        DB::beginTransaction();

        try {
            app("db")->getPdo()->exec("
            INSERT INTO `countries` (`name`, `code`, `currency_code`) VALUES ('Nigeria', 'NGA', 'NGN');
            INSERT INTO `countries` (`name`, `code`, `currency_code`) VALUES ('Kenya', 'KEN', 'KES');
            INSERT INTO `countries` (`name`, `code`, `currency_code`) VALUES ('Ghana', 'GHA', 'GHS');
            INSERT INTO `countries` (`name`, `code`, `currency_code`) VALUES ('Default', 'DEFAULT', 'EUR');

            -- 1
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('NGA', '0', 'Level 0');
            -- 2
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('NGA', '1', 'Level 1');
            -- 3
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('NGA', '2', 'Level 2');
            -- 4
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('NGA', '3', 'Level 3');
            -- 5
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('NGA', '4', 'Agent');
            -- 6
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('NGA', '5', 'Merchant');

            -- 7
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('KEN', '0', 'Entry');
            -- 8
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('KEN', '1', 'Individual');
            -- 9
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('KEN', '2', 'Merchant');
            -- 10
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('KEN', '3', 'Agents');

            -- 11
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('GHA', '0', 'Entry');
            -- 12
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('GHA', '1', 'Minimum');
            -- 13
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('GHA', '2', 'Medium');
            -- 14
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('GHA', '3', 'Enhanced');
            -- 15
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('GHA', '4', 'Agent');
            -- 16
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('GHA', '5', 'Merchant');

            -- 17
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('DEFAULT', '0', 'Level 0');
            -- 18
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('DEFAULT', '1', 'Level 1');
            -- 19
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('DEFAULT', '2', 'Level 2');
            -- 20
            INSERT INTO `tiers` (`country_code`, `level`, `name`) VALUES ('DEFAULT', '3', 'Level 3');

            -- limits for Nigeria
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('1', '0', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('1', '0', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('1', '0', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('1', '0', 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('1', '0', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('2', '10', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('2', '100', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('2', NULL, 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('2', '20', 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('2', '200', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('3', '50', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('3', '500', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('3', '100', 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('3', '200', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('4', '220', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('4', '220', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('4', '100', 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('4', '2200', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('5', '2200', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('5', '10000', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('5', '10000', 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('5', '10000', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('6', NULL, 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('6', NULL, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('6', NULL, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('6', NULL, 'max_total_balance');

            -- limits for Kenya
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('7', '0', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('7', '0', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('7', '0', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('7', '0', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('7', NULL, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('8', '2500', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('8', '2500', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('8', '700000', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('8', '1200', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('8', NULL, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('9', NULL, 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('9', NULL, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('9', NULL, 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('9', NULL, 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('10', NULL, 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('10', NULL, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('10', NULL, 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('10', NULL, 'max_debit_per_transfer');

            -- limits for Ghana
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('11', '0', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('11',  NULL, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('11', '0', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('11', '0', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('11', '0', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('12', '10', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('12',  NULL, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('12', '50', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('12', '150', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('12', '450', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('13', '50', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('13', '300', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('13', '1500', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('13', '3000', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('14', '200', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('14', '750', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('14', '3000', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('14', '7500', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('15', NULL, 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('15', NULL, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('15', NULL, 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('15', NULL, 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('16', NULL, 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('16', NULL, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('16', NULL, 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('16', NULL, 'max_total_debit_per_month');

            -- requirements Nigeria
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('1', 'Email or Mobile Number Verification on Sign Up', 'email_or_phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('2', 'Email', 'email');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('2', 'Phone', 'phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('2', 'Mother\'s Maiden Name', 'mother_maiden_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('2', 'Selfie Photography taken with phone\'s camera', 'selfie_photo');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('3', 'Bank Verification Number Validation', 'bank_number');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('3', 'Date of Birth', 'date_birth');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('3', 'Government Issued Identification Document', 'identification_document');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('4', 'Address Verification', 'address');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('5', 'Business/Corporate registration documents', 'business_registration_document');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('5', 'Verified Six Months\' Business Bank Account Statement', 'six_months_statement');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('5', 'Bank Guaranty', 'bank_guaranty');

            -- limits for Unknown (Default)
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('17', '0', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('17', '0', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('17', '0', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('17', '0', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('17', NULL, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('18', '10000', 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('18', '3000', 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('18', '10000', 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('18', '1000', 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('18', 5000, 'max_credit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('19', 50000, 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('19', 15000, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('19', 50000, 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('19', NULL, 'max_debit_per_transfer');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('20', NULL, 'max_total_balance');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('20', NULL, 'max_total_debit_per_day');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('20', NULL, 'max_total_debit_per_month');
            INSERT INTO `tier_limitations` (`tier_id`, `value`, `index`) VALUES ('20', NULL, 'max_debit_per_transfer');

            -- requirements Kenya
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('7', 'Email or Mobile Number Verification on Sign Up', 'email_or_phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Full Name', 'full_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Date of Birth', 'date_birth');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Email', 'email');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Phone', 'phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Mother\'s Maiden Name', 'mother_maiden_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Selfie Photography taken with phone\'s camera', 'selfie_photo');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('8', 'Government Issued Identification Document', 'identification_document');

            -- requirements Ghana
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('11', 'Email or Mobile Number Verification on Sign Up', 'email_or_phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('12', 'Full Name', 'full_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('12', 'Date of Birth', 'date_birth');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('12', 'Email', 'email');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('12', 'Phone', 'phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('12', 'Mother\'s Maiden Name', 'mother_maiden_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('12', 'Selfie Photography taken with phone\'s camera', 'selfie_photo');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('13', 'Government Issued Identification Document', 'identification_document');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('14', 'Address Verification', 'address');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('15', 'Business/Corporate registration documents', 'business_registration_document');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('15', 'Copies of IDs of the Directors and Beneficial Owners', 'directors_and_beneficial');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('15', 'Shareholders Documents and their nationalities', 'shareholders_documents');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('15', 'Income Tax Certificates', 'income_tax_certificates');

            -- requirements Unknown (default)
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('17', 'Email or Mobile Number Verification on Sign Up', 'email_or_phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Full Name', 'full_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Date of Birth', 'date_birth');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Email', 'email');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Phone', 'phone');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Mother\'s Maiden Name', 'mother_maiden_name');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Selfie Photography taken with phone\'s camera', 'selfie_photo');
            INSERT INTO `tier_requirements` (`tier_id`, `name`, `form_index`) VALUES ('18', 'Government Issued Identification Document', 'identification_document');
            ");
        } catch (\Throwable $e) {
            DB::rollBack();
            throw $e;
        }

        DB::commit();
    }
}
