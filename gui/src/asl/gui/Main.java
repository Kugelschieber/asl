/*
 * The MIT License
 *
 * Copyright 2015 Ozan Egitmen.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package asl.gui;

import java.awt.Color;
import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.prefs.Preferences;
import javax.swing.BorderFactory;
import javax.swing.JFileChooser;
import javax.swing.JTextField;
import javax.swing.SwingUtilities;
import javax.swing.UIManager;
import javax.swing.UnsupportedLookAndFeelException;
import javax.swing.border.Border;
import javax.swing.filechooser.FileNameExtensionFilter;
import javax.swing.plaf.ColorUIResource;

public class Main extends javax.swing.JFrame {

    Preferences prefs = Preferences.userRoot().node(this.getClass().getName());
    boolean inputError = false, outputError = false, aslError = false;

    public Main() {
        initComponents();
        getContentPane().setBackground(Color.WHITE);
        lblASLError.setText(" ");
        lblInputError.setText(" ");
        lblOutputError.setText(" ");
        txtASLDir.setText(prefs.get("aslDir", ""));
        txtInputDir.setText(prefs.get("inputDir", ""));
        txtOutputDir.setText(prefs.get("outputDir", ""));
        cbCompileAll.setSelected(prefs.getBoolean("compileAll", false));
        cbPrettyPrinting.setSelected(prefs.getBoolean("prettyPrinting", false));
    }

    private String fileChooser(String title, int fileType) {
        JFileChooser chooser = new JFileChooser();
        if (fileType == 0) {
            chooser.setFileFilter(new FileNameExtensionFilter("Executable", "exe"));
            chooser.setAcceptAllFileFilterUsed(false);
        }
        chooser.setFileSelectionMode(fileType);
        chooser.setDialogTitle(title);
        String selectedPath = "";
        if (chooser.showOpenDialog(null) == 0)
            selectedPath = chooser.getSelectedFile().toString();
        else
            chooser.cancelSelection();
        return selectedPath;
    }

    private void setErrorCondition(int i, boolean j) {
        if (i == 0) {
            lblASLError.setText(j ? "asl.exe isn't in this location! You can click this message to download it." : " ");
            aslError = j;
        } else if (i == 1) {
            lblInputError.setText(j ? "This folder doesn't exist!" : " ");
            inputError = j;
        } else {
            lblOutputError.setText(j ? "Output folder doesn't exsist! Click this message to create it." : " ");
            outputError = j;
        }
    }

    private void initComponents() {

        lblInput = new javax.swing.JLabel();
        txtInputDir = new javax.swing.JTextField();
        lblOutput = new javax.swing.JLabel();
        txtOutputDir = new javax.swing.JTextField();
        btnInput = new javax.swing.JButton();
        btnOutput = new javax.swing.JButton();
        lblASL = new javax.swing.JLabel();
        txtASLDir = new javax.swing.JTextField();
        btnASL = new javax.swing.JButton();
        jSeparator = new javax.swing.JSeparator();
        lblASLSmall = new javax.swing.JLabel();
        lblInputSmall = new javax.swing.JLabel();
        lblOutputSmall = new javax.swing.JLabel();
        cbCompileAll = new javax.swing.JCheckBox();
        cbPrettyPrinting = new javax.swing.JCheckBox();
        btnCompile = new javax.swing.JButton();
        lblASLError = new javax.swing.JLabel();
        lblInputError = new javax.swing.JLabel();
        lblOutputError = new javax.swing.JLabel();

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        setTitle("ASL GUI");
        setResizable(false);

        lblInput.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16));
        lblInput.setText("Input Directory:");
        lblInput.setOpaque(true);

        txtInputDir.setFont(new java.awt.Font("Segoe UI Light", 0, 16));

        lblOutput.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16));
        lblOutput.setText("Output Directory:");
        lblOutput.setOpaque(true);

        txtOutputDir.setFont(new java.awt.Font("Segoe UI Light", 0, 16));

        btnInput.setText("...");
        btnInput.setToolTipText("Opens a dialog to select input file");
        btnInput.setFocusable(false);
        btnInput.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnInputMouseClicked(evt);
            }
        });

        btnOutput.setText("...");
        btnOutput.setToolTipText("Opens a dialog to select output directory");
        btnOutput.setFocusable(false);
        btnOutput.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnOutputMouseClicked(evt);
            }
        });

        lblASL.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16));
        lblASL.setText("ASL Compiler Directory");
        lblASL.setOpaque(true);

        txtASLDir.setFont(new java.awt.Font("Segoe UI Light", 0, 16));

        btnASL.setText("...");
        btnASL.setToolTipText("Opens a dialog to select the compiler location");
        btnASL.setFocusable(false);
        btnASL.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnASLMouseClicked(evt);
            }
        });

        jSeparator.setToolTipText("");

        lblASLSmall.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10));
        lblASLSmall.setText("Location of the asl.exe file.");
        lblASLSmall.setOpaque(true);

        lblInputSmall.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10));
        lblInputSmall.setText("Directory of scripts that will be compiled in to the output directory.");
        lblInputSmall.setOpaque(true);

        lblOutputSmall.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10));
        lblOutputSmall.setText("Directory that the compiled .sqf script(s) will be saved in.");
        lblOutputSmall.setOpaque(true);

        cbCompileAll.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 11));
        cbCompileAll.setText("Compile all scripts in subfolders too.");
        cbCompileAll.setFocusable(false);
        cbCompileAll.addChangeListener(new javax.swing.event.ChangeListener() {
            public void stateChanged(javax.swing.event.ChangeEvent evt) {
                cbCompileAllStateChanged(evt);
            }
        });

        cbPrettyPrinting.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 11));
        cbPrettyPrinting.setText("Activate pretty printing.");
        cbPrettyPrinting.setFocusable(false);
        cbPrettyPrinting.addChangeListener(new javax.swing.event.ChangeListener() {
            public void stateChanged(javax.swing.event.ChangeEvent evt) {
                cbPrettyPrintingStateChanged(evt);
            }
        });

        btnCompile.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16));
        btnCompile.setText("Compile");
        btnCompile.setToolTipText("Opens a dialog to select output directory");
        btnCompile.setFocusable(false);
        btnCompile.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnCompileMouseClicked(evt);
            }
        });

        lblASLError.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10));
        lblASLError.setForeground(java.awt.Color.red);
        lblASLError.setText("Some error");
        lblASLError.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                lblASLErrorMouseClicked(evt);
            }
        });

        lblInputError.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10));
        lblInputError.setForeground(java.awt.Color.red);
        lblInputError.setText("Some error");

        lblOutputError.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10));
        lblOutputError.setForeground(java.awt.Color.red);
        lblOutputError.setText("Some error");
        lblOutputError.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                lblOutputErrorMouseClicked(evt);
            }
        });

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(getContentPane());
        getContentPane().setLayout(layout);
        layout.setHorizontalGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(jSeparator).addGroup(layout.createSequentialGroup().addGap(15, 15, 15).addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(lblOutputError, javax.swing.GroupLayout.PREFERRED_SIZE, 371, javax.swing.GroupLayout.PREFERRED_SIZE).addComponent(lblInputError, javax.swing.GroupLayout.PREFERRED_SIZE, 371, javax.swing.GroupLayout.PREFERRED_SIZE).addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.TRAILING, false).addComponent(lblASLError, javax.swing.GroupLayout.Alignment.LEADING, javax.swing.GroupLayout.PREFERRED_SIZE, 371, javax.swing.GroupLayout.PREFERRED_SIZE).addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(lblASL).addComponent(lblASLSmall).addGroup(layout.createSequentialGroup().addComponent(txtASLDir, javax.swing.GroupLayout.PREFERRED_SIZE, 320, javax.swing.GroupLayout.PREFERRED_SIZE).addGap(6, 6, 6).addComponent(btnASL)).addGroup(layout.createSequentialGroup().addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.TRAILING).addComponent(txtOutputDir, javax.swing.GroupLayout.Alignment.LEADING, javax.swing.GroupLayout.PREFERRED_SIZE, 320, javax.swing.GroupLayout.PREFERRED_SIZE).addComponent(lblOutput, javax.swing.GroupLayout.Alignment.LEADING)).addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED).addComponent(btnOutput)).addComponent(lblOutputSmall).addComponent(lblInput).addGroup(layout.createSequentialGroup().addComponent(txtInputDir, javax.swing.GroupLayout.PREFERRED_SIZE, 320, javax.swing.GroupLayout.PREFERRED_SIZE).addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED).addComponent(btnInput)).addComponent(lblInputSmall).addComponent(cbCompileAll).addComponent(cbPrettyPrinting))).addComponent(btnCompile, javax.swing.GroupLayout.PREFERRED_SIZE, 373, javax.swing.GroupLayout.PREFERRED_SIZE)).addGap(20, 20, 20)));
        layout.setVerticalGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addGroup(layout.createSequentialGroup().addGap(6, 6, 6).addComponent(lblASL).addGap(3, 3, 3).addComponent(lblASLSmall).addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED).addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(txtASLDir, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE).addComponent(btnASL, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE)).addGap(4, 4, 4).addComponent(lblASLError).addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED).addComponent(jSeparator, javax.swing.GroupLayout.PREFERRED_SIZE, 10, javax.swing.GroupLayout.PREFERRED_SIZE).addGap(3, 3, 3).addComponent(lblInput).addGap(3, 3, 3).addComponent(lblInputSmall, javax.swing.GroupLayout.PREFERRED_SIZE, 14, javax.swing.GroupLayout.PREFERRED_SIZE).addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED).addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(txtInputDir, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE).addComponent(btnInput, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE)).addGap(4, 4, 4).addComponent(lblInputError).addGap(6, 6, 6).addComponent(lblOutput).addGap(3, 3, 3).addComponent(lblOutputSmall, javax.swing.GroupLayout.PREFERRED_SIZE, 14, javax.swing.GroupLayout.PREFERRED_SIZE).addGap(6, 6, 6).addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(txtOutputDir, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE).addComponent(btnOutput, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE)).addGap(4, 4, 4).addComponent(lblOutputError).addGap(6, 6, 6).addComponent(cbCompileAll).addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.UNRELATED).addComponent(cbPrettyPrinting).addGap(11, 11, 11).addComponent(btnCompile, javax.swing.GroupLayout.PREFERRED_SIZE, 41, javax.swing.GroupLayout.PREFERRED_SIZE).addGap(11, 11, 11)));

        pack();
        setLocationRelativeTo(null);
    }

    private void btnInputMouseClicked(java.awt.event.MouseEvent evt) {
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String path = fileChooser("Select input directory", 1);
            File inputDir = new File(path);
            if (inputDir.exists()) {
                prefs.put("inputDir", path);
                txtInputDir.setText(path);
                if (inputError)
                    setErrorCondition(1, false);
            }
        }
    }

    private void btnOutputMouseClicked(java.awt.event.MouseEvent evt) {
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String path = fileChooser("Select output directory", 1);
            File outputDir = new File(path);
            if (outputDir.exists()) {
                prefs.put("outputDir", path);
                txtOutputDir.setText(path);
                if (outputError)
                    setErrorCondition(2, false);
            }
        }
    }

    private void btnASLMouseClicked(java.awt.event.MouseEvent evt) {
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String path = fileChooser("Select 'asl.exe' location", 0);
            File asl = new File(path);
            if (asl.exists()) {
                prefs.put("aslDir", path);
                txtASLDir.setText(path);
                if (aslError)
                    setErrorCondition(0, false);
            }
        }
    }

    private void btnCompileMouseClicked(java.awt.event.MouseEvent evt) {
        if (SwingUtilities.isLeftMouseButton(evt)) {
            JTextField[] dirFields = {txtASLDir, txtInputDir, txtOutputDir};
            for (byte i = 0; i < 3; i++) {
                File bleh = new File(dirFields[i].getText());
                setErrorCondition(i, !bleh.exists());
            }
            if (aslError || inputError || outputError)
                return;
            String prettyPrinting = cbPrettyPrinting.isSelected() ? "-pretty" : "", compileAll = cbCompileAll.isSelected() ? "-r" : "", asl = txtASLDir.getText(), input = txtInputDir.getText(), output = txtOutputDir.getText(), error = " ";    
            try {
                Process aslProcess = new ProcessBuilder(asl, compileAll, prettyPrinting, input, output).start();
                BufferedReader br = new BufferedReader(new InputStreamReader(aslProcess.getInputStream()));
                String line;
                while ((line = br.readLine()) != null) {
                    if (line.toLowerCase().contains("panic")) {
                        error = line;
                    }
                }
                aslProcess.waitFor();
            } catch (IOException | InterruptedException ex) {
                Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
            }
            if (!error.equals(" ")) {
                DlgError showError = new DlgError(this, true, error);
                showError.setLocationRelativeTo(this);
                showError.setVisible(true);
            }
        }
    }

    private void cbCompileAllStateChanged(javax.swing.event.ChangeEvent evt) {
        prefs.putBoolean("compileAll", cbCompileAll.isSelected());
    }

    private void cbPrettyPrintingStateChanged(javax.swing.event.ChangeEvent evt) {
        prefs.putBoolean("prettyPrinting", cbPrettyPrinting.isSelected());
    }

    private void lblOutputErrorMouseClicked(java.awt.event.MouseEvent evt) {
        if (SwingUtilities.isLeftMouseButton(evt) && outputError) {
            new File(txtOutputDir.getText()).mkdirs();
            setErrorCondition(2, false);
        }
    }

    private void lblASLErrorMouseClicked(java.awt.event.MouseEvent evt) {
        if (SwingUtilities.isLeftMouseButton(evt) && aslError) {
            try {
                URI github = new URI("https://github.com/DeKugelschieber/asl/releases");
                java.awt.Desktop.getDesktop().browse(github);
            } catch (URISyntaxException | IOException ex) {
                Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
            }
            setErrorCondition(0, false);
        }
    }

    public static void main(String args[]) {
        try {
            UIManager.setLookAndFeel("com.sun.java.swing.plaf.windows.WindowsLookAndFeel");
        } catch (ClassNotFoundException | InstantiationException | IllegalAccessException | UnsupportedLookAndFeelException ex) {
            Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
        }
        UIManager.put("ToolTip.background", new ColorUIResource(255, 255, 255));
        UIManager.put("ToolTip.foreground", new ColorUIResource(87, 87, 87));
        Border lineBorder = BorderFactory.createLineBorder(new Color(118, 118, 118));
        UIManager.put("ToolTip.border", lineBorder);
        Border compoundBorder = BorderFactory.createCompoundBorder(UIManager.getBorder("ToolTip.border"), BorderFactory.createEmptyBorder(0, 2, 2, 3));
        UIManager.put("ToolTip.border", compoundBorder);
        java.awt.EventQueue.invokeLater(() -> {
            new Main().setVisible(true);
        });
    }

    private javax.swing.JButton btnASL;
    private javax.swing.JButton btnCompile;
    private javax.swing.JButton btnInput;
    private javax.swing.JButton btnOutput;
    private javax.swing.JCheckBox cbCompileAll;
    private javax.swing.JCheckBox cbPrettyPrinting;
    private javax.swing.JSeparator jSeparator;
    private javax.swing.JLabel lblASL;
    private javax.swing.JLabel lblASLError;
    private javax.swing.JLabel lblASLSmall;
    private javax.swing.JLabel lblInput;
    private javax.swing.JLabel lblInputError;
    private javax.swing.JLabel lblInputSmall;
    private javax.swing.JLabel lblOutput;
    private javax.swing.JLabel lblOutputError;
    private javax.swing.JLabel lblOutputSmall;
    private javax.swing.JTextField txtASLDir;
    private javax.swing.JTextField txtInputDir;
    private javax.swing.JTextField txtOutputDir;
}
