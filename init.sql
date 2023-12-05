-- Create the Doctor table
CREATE TABLE IF NOT EXISTS doctor (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL
);

-- Create the DoctorSlots table
CREATE TABLE IF NOT EXISTS doctor_slots (
    id SERIAL PRIMARY KEY,
    doctor_id INT,
    doctor_email VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    hour VARCHAR(255) NOT NULL,
    is_booked BOOLEAN DEFAULT false
);

-- Create the Appointment table
CREATE TABLE IF NOT EXISTS appointment (
    id SERIAL PRIMARY KEY,
    patient_id INT,
    doctor_email VARCHAR(255) NOT NULL,
    patient_email VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    hour VARCHAR(255) NOT NULL
);

-- Create the Patient table
CREATE TABLE IF NOT EXISTS patient (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL
);
