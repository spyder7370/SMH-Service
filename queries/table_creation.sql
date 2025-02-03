
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists departments (
	id uuid primary key default uuid_generate_v4(),
	name text not null,
	description text
)
insert into departments (name) values 'Emergency Department (ER)',
'Cardiology'
'Neurology'
'Pediatrics',
'Oncology',
'Orthopedics',
'Gynecology & Obstetrics',
'Urology',
'Nephrology',
'Gastroenterology',
'Pulmonology',
'Endocrinology',
'Rheumatology',
'Dermatology',
'Ophthalmology',
'Otolaryngology',
'Hematology',
'Psychiatry & Mental Health',
'Radiology',
'Anesthesiology',
'Intensive Care Unit',
'Neonatal Intensive Care Unit',
'Geriatrics',
'Physical Therapy & Rehabilitation',
'Infectious Diseases',
'Pathology',


create table if not exists doctors (
	id uuid primary key default uuid_generate_v4(),
	name text not null,
    password text not null,
	specialisation text not null,
	phone text,
	email text,
	dob date,
    address text
)

create index idx_specialisations on doctors (specialisation)

create table if not exist doctor_timetable (
    audit_id uuid primary key default uuid_generate_v4(),
    doctor_id uuid not null references doctors(id),
    day_of_week enum('MONDAY','TUESDAY','WEDNESDAY','THRUSDAY','FRIDAY','SATURDAY','SUNDAY') not null,
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    effective_from timestamp with time zone not null,
    effective_to timestamp with time zone not null
)

create index idx_doctor_timetable on doctor_timetable (effective_from desc, doctor_id) 
create table if not exist doctors_current_schedule (
    slot_id uuid primary key default uuid_generate_v4(),
    doctor_id uuid not null references doctors(id),
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
)

create index idx_doctor_latest_schedule on doctors_current_schedule (start_time desc, doctor_id)



create table if not exists doctor_departments (
    doctor_id uuid not null references doctors(id) on delete cascade,
    department_id uuid not null references departments(id) on delete cascade,
    assigned_date date not null,
    role text,
    primary key(doctor_id,department_id)
)

create table if not exists patients (
    id uuid primary key default uuid_generate_v4(),
    name text not null,
    password text not null,
    dob date,
    phone text,
    email text,
    address text
)

--drop table appointments
create table if not exists appointments (
    id uuid primary key default uuid_generate_v4(),
    doctor_id uuid not null references doctors(id),
    patient_id uuid not null references patients(id),
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    notes text,
    status text,
    constraint chk_time_validity check (end_time > start_time)
)

create index idx_appointments on appointments (start_time desc, doctor_id , patient_id)


create table if not exist manufactures (
    id uuid primary key default uuid_generate_v4(),
    name text not null,
    phone text,
    email text,
    address text
)

create table if not exist suppliers (
    id uuid primary key default uuid_generate_v4(),
    name text not null,
    phone text,
    email text,
    address text
)

create table if not exist medicine (
    id uuid primary key default uuid_generate_v4(),
    name text not null,
    description text
)

create table if not exist medicine_purchases (
    purchase_id uuid primary key 
)
